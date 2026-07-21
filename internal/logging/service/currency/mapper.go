package currency

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func MappingStarted(operation string) func() { return log.MappingStarted("currency", operation) }
func NilInput(operation string)              { log.NilInput("currency", operation) }
