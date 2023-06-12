package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/handler"
	"github.com/hifat/sodium-api/internal/middleware"
	_ "github.com/joho/godotenv/autoload"
)

type routes struct {
	router     *gin.RouterGroup
	handler    handler.Handler
	middleware middleware.Middleware
}

func New(router *gin.RouterGroup, handler handler.Handler, middleware middleware.Middleware) *routes {
	return &routes{
		router,
		handler,
		middleware,
	}
}

func (r routes) Register() {
	r.authRoute()
}
