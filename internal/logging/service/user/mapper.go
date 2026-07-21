package user

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func MappingStarted(operation string) func() { return log.MappingStarted("user", operation) }
func NilInput(operation string)              { log.NilInput("user", operation) }
