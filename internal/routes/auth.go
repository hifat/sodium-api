package routes

import (
	"github.com/hifat/sodium-api/internal/handler/authHandler"
	"github.com/hifat/sodium-api/internal/middleware"
	"github.com/hifat/sodium-api/internal/repository/authRepo"
	"github.com/hifat/sodium-api/internal/repository/userRepo"
	"github.com/hifat/sodium-api/internal/service/authService"
	"github.com/hifat/sodium-api/internal/service/middlewareService"
)

func (r routes) authRoute() {
	newAuthRepo := authRepo.NewAuthRepository(r.orm)
	newUserRepo := userRepo.NewUserRepository(r.orm)

	newAuthService := authService.NewAuthService(newAuthRepo, newUserRepo)
	newAuthMiddlewareService := middlewareService.NewAuthMiddlewareService(newAuthRepo)

	newAuthHandler := authHandler.NewAuthHandler(newAuthService)

	newAuthMiddleware := middleware.NewAuthMiddleware(newAuthMiddlewareService)

	authRoute := r.router.Group("/auth")
	{
		authRoute.POST("/register", newAuthHandler.Register)
		authRoute.POST("/login", newAuthHandler.Login)
		authRoute.POST("/token/refresh", newAuthMiddleware.AuthRefreshGuard, newAuthHandler.CreateRefreshToken)
	}
}
