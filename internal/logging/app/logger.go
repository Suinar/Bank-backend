package app

import (
	stdlog "log"
	"time"
)

func Configure() {
	stdlog.SetFlags(stdlog.Ldate | stdlog.Ltime | stdlog.Lmicroseconds | stdlog.LUTC)
}

func Started() { stdlog.Printf("application: starting") }

func Stopped() { stdlog.Printf("application: stopped") }

func ConfigurationLoading() { stdlog.Printf("application: loading configuration") }

func ConfigurationFailed(err error) {
	stdlog.Printf("application: configuration failed error=%q", err)
}

func ConfigurationLoaded(environment, httpAddress, repositoryGRPC, exchangeRateGRPC string) {
	stdlog.Printf(
		"application: configuration loaded environment=%s http_address=%s repository_grpc=%s exchange_rate_grpc=%s",
		environment, httpAddress, repositoryGRPC, exchangeRateGRPC,
	)
}

func GRPCClientsInitializing() { stdlog.Printf("application: initializing gRPC clients") }

func GRPCClientsInitializationFailed(err error) {
	stdlog.Printf("application: gRPC client initialization failed error=%q", err)
}

func GRPCClientsInitialized() { stdlog.Printf("application: gRPC clients initialized") }

func GRPCClientsClosing() { stdlog.Printf("application: closing gRPC clients") }

func GRPCClientsClosingFailed(err error) {
	stdlog.Printf("application: closing gRPC clients failed error=%q", err)
}

func GRPCClientsClosed() { stdlog.Printf("application: gRPC clients closed") }

func ComponentsInitializing() {
	stdlog.Printf("application: initializing services and HTTP handlers")
}

func ComponentsInitialized() {
	stdlog.Printf("application: services and HTTP handlers initialized")
}

func HTTPServerStarting(address string) {
	stdlog.Printf("application: HTTP server starting address=%s", address)
}

func ShutdownSignalReceived() { stdlog.Printf("application: shutdown signal received") }

func HTTPServerShuttingDown(timeout time.Duration) {
	stdlog.Printf("application: shutting down HTTP server timeout=%s", timeout)
}

func HTTPServerShutdownFailed(err error) {
	stdlog.Printf("application: HTTP server shutdown failed error=%q", err)
}

func HTTPServerFailed(err error) {
	stdlog.Printf("application: HTTP server failed error=%q", err)
}

func HTTPServerStopped() { stdlog.Printf("application: HTTP server stopped") }
