package credit

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func MappingStarted(operation string) func() { return log.MappingStarted("credit", operation) }
func NilInput(operation string)              { log.NilInput("credit", operation) }
