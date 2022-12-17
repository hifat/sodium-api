package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/api/rest"
	"github.com/hifat/hifat-blog-api/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	if os.Getenv("APP_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	db, err := database.PostgresDB()
	if err != nil {
		panic(err)
	}
	_ = db

	rest.API()
}
