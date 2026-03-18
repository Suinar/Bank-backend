package core

import "errors"

var (
	NotFound            = errors.New("not found")
	BadRequest          = errors.New("bad request")
	InternalServerError = errors.New("internal server error")
)
