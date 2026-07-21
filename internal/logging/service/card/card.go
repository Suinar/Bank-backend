package card

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func OperationStarted(operation string) func()  { return log.OperationStarted("card", operation) }
func ValidationFailed(operation, reason string) { log.ValidationFailed("card", operation, reason) }
func DependencyFailed(operation, dependency string, err error) {
	log.DependencyFailed("card", operation, dependency, err)
}
