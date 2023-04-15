package rest

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/docs"
	"github.com/hifat/hifat-blog-api/internal/routes"
	"github.com/hifat/hifat-blog-api/internal/utils/validity"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func API() {
	/* ---------------------------- Validator config ---------------------------- */

	validity.Register()

	/* ------------------------------- Swag config ------------------------------ */
	
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Sodium API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	/* --------------------------- Running API server --------------------------- */
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	routes.AuthRoute(api)

	router.Run(fmt.Sprintf("%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	))
}
