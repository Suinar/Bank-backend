package common

import (
	"errors"

	"google.golang.org/grpc/status"
)

// IsError matches local domain errors and errors transported by legacy gRPC
// handlers that currently return the domain message with status Unknown.
func IsError(err, target error) bool {
	if errors.Is(err, target) {
		return true
	}
	if err == nil || target == nil {
		return false
	}
	return status.Convert(err).Message() == target.Error()
}
