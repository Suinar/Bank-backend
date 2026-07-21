package exhange_rate

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func OperationStarted(operation string) func() {
	return log.OperationStarted("exchange_rate", operation)
}
func ValidationFailed(operation, reason string) {
	log.ValidationFailed("exchange_rate", operation, reason)
}
func DependencyFailed(operation, dependency string, err error) {
	log.DependencyFailed("exchange_rate", operation, dependency, err)
}
