package common

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestIsError(t *testing.T) {
	target := errors.New("not found")
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{name: "domain error", err: target, want: true},
		{name: "legacy grpc error", err: status.Error(codes.Unknown, "not found"), want: true},
		{name: "different grpc error", err: status.Error(codes.Unknown, "internal"), want: false},
		{name: "nil", err: nil, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsError(test.err, target); got != test.want {
				t.Fatalf("IsError() = %t, want %t", got, test.want)
			}
		})
	}
}
