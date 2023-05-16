package rest

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/docs"
	"github.com/hifat/sodium-api/internal/database"
	"github.com/hifat/sodium-api/internal/routes"
	"github.com/hifat/sodium-api/internal/utils/validity"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func API() {
	/* --------------------------------- Init DB -------------------------------- */
	orm := database.PostgresDB()
	db, err := orm.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

	r := routes.New(orm, api)
	r.Register()

	router.Run(fmt.Sprintf("%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	))
}
