package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/handler"
	_ "github.com/joho/godotenv/autoload"
)

type routes struct {
	router *gin.RouterGroup
	h      handler.Handler
}

func New(router *gin.RouterGroup, h handler.Handler) *routes {
	return &routes{
		router,
		h,
	}
}

func (r routes) Register() {
	r.authRoute()
}
