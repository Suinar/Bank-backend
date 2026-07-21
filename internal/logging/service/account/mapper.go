package account

import (
	stdlog "log"
	"time"
)

func MappingStarted(operation string) func() {
	startedAt := time.Now()
	stdlog.Printf("account mapper: mapping started operation=%s", operation)

	return func() {
		stdlog.Printf(
			"account mapper: mapping completed operation=%s duration=%s",
			operation,
			time.Since(startedAt),
		)
	}
}

func NilInput(operation string) {
	stdlog.Printf(
		"account mapper: nil input operation=%s",
		operation,
	)
}
