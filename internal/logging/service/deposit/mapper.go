package deposit

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func MappingStarted(operation string) func() { return log.MappingStarted("deposit", operation) }
func NilInput(operation string)              { log.NilInput("deposit", operation) }
