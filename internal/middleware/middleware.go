package middleware

import (
	"github.com/google/wire"
)

var MiddlewareSet = wire.NewSet(NewMiddleware)

type Middleware struct {
	AuthMiddleware
}

func NewMiddleware(ah AuthMiddleware) Middleware {
	return Middleware{
		AuthMiddleware: ah,
	}
}
