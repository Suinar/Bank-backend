package client

import (
	stdlog "log"
	"time"
)

func ConnectionStarting(target string, timeout time.Duration) {
	stdlog.Printf("gRPC client: connection starting target=%s timeout=%s", target, timeout)
}

func ConnectionEstablished(target string, duration time.Duration) {
	stdlog.Printf("gRPC client: connection established target=%s duration=%s", target, duration)
}

func ConnectionFailed(target string, duration time.Duration, err error) {
	stdlog.Printf("gRPC client: connection failed target=%s duration=%s error=%q", target, duration, err)
}

func Initializing(name, target string) {
	stdlog.Printf("gRPC client: initializing name=%s target=%s", name, target)
}

func Initialized(name, target string) {
	stdlog.Printf("gRPC client: initialized name=%s target=%s", name, target)
}

func InitializationFailed(name, target string, err error) {
	stdlog.Printf("gRPC client: initialization failed name=%s target=%s error=%q", name, target, err)
}

func ContainerInitializing() { stdlog.Printf("gRPC clients: container initializing") }

func ContainerInitialized() { stdlog.Printf("gRPC clients: container initialized") }

func RollbackStarting(name string) {
	stdlog.Printf("gRPC clients: rolling back connection name=%s", name)
}

func RollbackFailed(name string, err error) {
	stdlog.Printf("gRPC clients: connection rollback failed name=%s error=%q", name, err)
}

func Closing(name string) { stdlog.Printf("gRPC client: connection closing name=%s", name) }

func Closed(name string) { stdlog.Printf("gRPC client: connection closed name=%s", name) }

func CloseFailed(name string, err error) {
	stdlog.Printf("gRPC client: connection close failed name=%s error=%q", name, err)
}

func AllClosed() { stdlog.Printf("gRPC clients: all connections closed") }
