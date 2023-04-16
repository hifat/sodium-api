package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/database"
	"github.com/hifat/sodium-api/internal/handler/authHandler"
	"github.com/hifat/sodium-api/internal/repository/authRepo"
	"github.com/hifat/sodium-api/internal/service/authService"
)

func AuthRoute(r *gin.RouterGroup) {
	newauthRepo := authRepo.NewauthRepository(database.PostgresDB())
	newAuthService := authService.NewAuthService(newauthRepo)
	newAuthHandler := authHandler.NewAuthHandler(newAuthService)

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", newAuthHandler.Register)
		authRoute.POST("/login", newAuthHandler.Login)
	}
}
