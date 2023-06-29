package di

import (
	"github.com/hifat/sodium-api/internal/handler"
	"github.com/hifat/sodium-api/internal/middleware"
)

type Adapter struct {
	Middleware middleware.Middleware
	Handler    handler.Handler
}

func NewAdapter(m middleware.Middleware, h handler.Handler) Adapter {
	return Adapter{
		Middleware: m,
		Handler:    h,
	}
}
