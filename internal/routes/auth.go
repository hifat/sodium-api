package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/database"
	"github.com/hifat/sodium-api/internal/handler/authHandler"
	"github.com/hifat/sodium-api/internal/repository/authRepo"
	"github.com/hifat/sodium-api/internal/repository/userRepo"
	"github.com/hifat/sodium-api/internal/service/authService"
)

func AuthRoute(r *gin.RouterGroup) {
	newAuthRepo := authRepo.NewauthRepository(database.PostgresDB())
	newUserRepo := userRepo.NewUserRepository(database.PostgresDB())

	newAuthService := authService.NewAuthService(newAuthRepo, newUserRepo)

	newAuthHandler := authHandler.NewAuthHandler(newAuthService)

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", newAuthHandler.Register)
		authRoute.POST("/login", newAuthHandler.Login)
		authRoute.GET("/token/refresh", newAuthHandler.CreateRefreshToken)
	}
}
