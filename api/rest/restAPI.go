package rest

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/internal/routes"
)

func API() {
	router := gin.Default()
	api := router.Group("/api")

	routes.AuthRoute(api)

	router.Run(fmt.Sprintf("%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	))
}
