package core

import "errors"

var (
	NotFound             = errors.New("not found")
	BadRequest           = errors.New("bad request")
	InternalServerError  = errors.New("internal server error")
	MoreThanOneCharacter = errors.New("more than one character")
	CacheGetError        = errors.New("redis HGetAll failed")
	CacheSetError        = errors.New("redis HSet failed")
	CacheDeleteError     = errors.New("redis HDel failed")
	CacheMapError        = errors.New("redis mapping failed")
)
