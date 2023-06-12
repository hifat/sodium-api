package handler

import (
	"github.com/google/wire"
	"github.com/hifat/sodium-api/internal/handler/authHandler"
)

var HandlerSet = wire.NewSet(NewHandler)

type Handler struct {
	AuthHandler authHandler.AuthHandler
}

func NewHandler(ah authHandler.AuthHandler) Handler {
	return Handler{
		AuthHandler: ah,
	}
}
