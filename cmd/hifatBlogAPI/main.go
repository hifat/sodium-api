package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/api/rest"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	if os.Getenv("APP_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	rest.API()
}
