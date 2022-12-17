package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/hifat-blog-api/internal/database"
	"github.com/hifat/hifat-blog-api/internal/handler/authHandler"
	"github.com/hifat/hifat-blog-api/internal/repository/authRepo"
	"github.com/hifat/hifat-blog-api/internal/service/authService"
)

func AuthRoute(r *gin.RouterGroup) {
	newAuthRepo := authRepo.NewAuthRepository(database.PostgresDB())
	newAuthService := authService.NewAuthService(newAuthRepo)
	newAuthHandler := authHandler.NewAuthHandler(newAuthService)

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", newAuthHandler.Register)
	}
}
