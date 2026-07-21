package user

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func OperationStarted(operation string) func()  { return log.OperationStarted("user", operation) }
func ValidationFailed(operation, reason string) { log.ValidationFailed("user", operation, reason) }
func DependencyFailed(operation, dependency string, err error) {
	log.DependencyFailed("user", operation, dependency, err)
}
