package card

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func MappingStarted(operation string) func() { return log.MappingStarted("card", operation) }
func NilInput(operation string)              { log.NilInput("card", operation) }
