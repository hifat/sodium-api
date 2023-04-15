package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/internal/database"
	"github.com/hifat/sodium-api/internal/handler/authHandler"
	"github.com/hifat/sodium-api/internal/repository/userRepo"
	"github.com/hifat/sodium-api/internal/service/authService"
)

func AuthRoute(r *gin.RouterGroup) {
	newUserRepo := userRepo.NewUserRepository(database.PostgresDB())
	newAuthService := authService.NewAuthService(newUserRepo)
	newAuthHandler := authHandler.NewAuthHandler(newAuthService)

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", newAuthHandler.Register)
	}
}
