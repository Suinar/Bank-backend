package credit

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func OperationStarted(operation string) func()  { return log.OperationStarted("credit", operation) }
func ValidationFailed(operation, reason string) { log.ValidationFailed("credit", operation, reason) }
func DependencyFailed(operation, dependency string, err error) {
	log.DependencyFailed("credit", operation, dependency, err)
}
