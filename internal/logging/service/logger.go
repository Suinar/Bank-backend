package service

import (
	stdlog "log"
	"time"
)

func OperationStarted(serviceName, operation string) func() {
	startedAt := time.Now()
	stdlog.Printf("%s service: operation started operation=%s", serviceName, operation)
	return func() {
		stdlog.Printf("%s service: operation completed operation=%s duration=%s", serviceName, operation, time.Since(startedAt))
	}
}

func MappingStarted(serviceName, operation string) func() {
	startedAt := time.Now()
	stdlog.Printf("%s mapper: mapping started operation=%s", serviceName, operation)
	return func() {
		stdlog.Printf("%s mapper: mapping completed operation=%s duration=%s", serviceName, operation, time.Since(startedAt))
	}
}

func NilInput(serviceName, operation string) {
	stdlog.Printf("%s mapper: nil input operation=%s", serviceName, operation)
}

func ValidationFailed(serviceName, operation, reason string) {
	stdlog.Printf("%s service: validation failed operation=%s reason=%s", serviceName, operation, reason)
}

func DependencyFailed(serviceName, operation, dependency string, err error) {
	stdlog.Printf("%s service: dependency failed operation=%s dependency=%s error=%q", serviceName, operation, dependency, err)
}
