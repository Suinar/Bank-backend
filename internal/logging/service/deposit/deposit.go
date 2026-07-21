package deposit

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func OperationStarted(operation string) func()  { return log.OperationStarted("deposit", operation) }
func ValidationFailed(operation, reason string) { log.ValidationFailed("deposit", operation, reason) }
func DependencyFailed(operation, dependency string, err error) {
	log.DependencyFailed("deposit", operation, dependency, err)
}
