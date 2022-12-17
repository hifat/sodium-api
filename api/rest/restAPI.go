package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func API() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})

	router.Run(fmt.Sprintf("%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	))
}
