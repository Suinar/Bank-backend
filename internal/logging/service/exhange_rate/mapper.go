package exhange_rate

import log "github.com/kVinsom/Bank-backend/internal/logging/service"

func MappingStarted(operation string) func() { return log.MappingStarted("exchange_rate", operation) }
func NilInput(operation string)              { log.NilInput("exchange_rate", operation) }
