package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/kVinsom/Bank-backend/docs"
	kafkaBroker "github.com/kVinsom/Bank-backend/internal/brokers/kafka"
	"github.com/kVinsom/Bank-backend/internal/configs"
	grpcClient "github.com/kVinsom/Bank-backend/internal/delivery/grps/client"
	httpHandler "github.com/kVinsom/Bank-backend/internal/delivery/http/handler"
	httpMiddleware "github.com/kVinsom/Bank-backend/internal/delivery/http/middleware"
	log "github.com/kVinsom/Bank-backend/internal/logging/app"
	serviceContainer "github.com/kVinsom/Bank-backend/internal/service"
)

// @title Bank Backend API
// @version 1.0
// @description HTTP API for users, accounts, cards, credits, deposits, currencies, and exchange rates.
// @BasePath /api/v1
// @schemes http https

const (
	// Explicit timeouts protect the server from slow or abandoned connections.
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 15 * time.Second
	writeTimeout      = 15 * time.Second
	idleTimeout       = 60 * time.Second
	shutdownTimeout   = 10 * time.Second
)

func main() {
	log.Configure()
	log.Started()
	defer log.Stopped()

	// Validate configuration before allocating network resources.
	log.ConfigurationLoading()
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.ConfigurationFailed(err)
		return
	}
	log.ConfigurationLoaded(
		cfg.AppConfig.Environment,
		net.JoinHostPort(cfg.HTTPConfig.Host, fmt.Sprint(cfg.HTTPConfig.Port)),
		cfg.GRPCConfig.RepositoryURL,
		cfg.GRPCConfig.ExchangeRateURL,
	)

	// SIGINT and SIGTERM initiate one shared graceful-shutdown path.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// The client container owns all outbound gRPC connections.
	log.GRPCClientsInitializing()
	clients, err := grpcClient.InitClients(ctx, *cfg)
	if err != nil {
		log.GRPCClientsInitializationFailed(err)
		return
	}
	log.GRPCClientsInitialized()
	defer func() {
		log.GRPCClientsClosing()
		if err := clients.Close(); err != nil {
			log.GRPCClientsClosingFailed(err)
			return
		}
		log.GRPCClientsClosed()
	}()

	var publisher *kafkaBroker.Publisher
	if cfg.KafkaConfig.Enabled {
		kafkaCtx, cancelKafka := context.WithTimeout(ctx, 60*time.Second)
		publisher, err = kafkaBroker.Open(kafkaCtx, cfg.KafkaConfig)
		cancelKafka()
		if err != nil {
			log.ConfigurationFailed(fmt.Errorf("initialize Kafka publisher: %w", err))
			return
		}
		defer func() {
			if err := publisher.Close(); err != nil {
				log.ConfigurationFailed(fmt.Errorf("close Kafka publisher: %w", err))
			}
		}()
	}

	// Wire dependencies from transport clients through services to HTTP handlers.
	log.ComponentsInitializing()
	services := serviceContainer.InitServices(cfg, clients.Repositories, clients.ExchangeRate)
	log.ComponentsInitialized()
	if publisher != nil {
		router := httpHandler.InitHandlers(services, httpMiddleware.KafkaAudit(publisher))
		runServer(ctx, cfg, router)
		return
	}
	router := httpHandler.InitHandlers(services)
	runServer(ctx, cfg, router)
}

func runServer(ctx context.Context, cfg *configs.Config, router http.Handler) {
	server := &http.Server{
		Addr:              net.JoinHostPort(cfg.HTTPConfig.Host, fmt.Sprint(cfg.HTTPConfig.Port)),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	// Buffering prevents the serving goroutine from blocking during shutdown.
	serverErrors := make(chan error, 1)
	go func() {
		log.HTTPServerStarting(server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Stop on either an operating-system signal or an unexpected server failure.
	select {
	case <-ctx.Done():
		log.ShutdownSignalReceived()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		log.HTTPServerShuttingDown(shutdownTimeout)
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.HTTPServerShutdownFailed(err)
			return
		}
		log.HTTPServerStopped()
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			log.HTTPServerFailed(err)
			return
		}
		log.HTTPServerStopped()
	}
}
