package client

import (
	"context"
	"time"

	configs "github.com/Suinar/Bank-backend/bankBackend/configs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	notification "github.com/Suinar/Bank-backend/bankBackend/proto/notification"
)

func NewNotificationClient(ctx context.Context, cfg configs.Config) (notification.NotificationServiceClient, *grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		cfg.Grpc.NotificationUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, nil, err
	}

	client := notification.NewNotificationServiceClient(conn)

	return client, conn, nil
}
