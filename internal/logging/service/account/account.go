package account

import (
	stdlog "log"
	"time"
)

func OperationStarted(operation string) func() {
	startedAt := time.Now()
	stdlog.Printf("account service: operation started operation=%s", operation)

	return func() {
		stdlog.Printf(
			"account service: operation completed operation=%s duration=%s",
			operation,
			time.Since(startedAt),
		)
	}
}

func ValidationFailed(operation, reason string) {
	stdlog.Printf(
		"account service: validation failed operation=%s reason=%s",
		operation,
		reason,
	)
}

func DependencyFailed(operation, dependency string, err error) {
	stdlog.Printf(
		"account service: dependency failed operation=%s dependency=%s error=%q",
		operation,
		dependency,
		err,
	)
}
