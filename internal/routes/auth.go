package routes

func (r routes) authRoute() {
	// newAuthRepo := authRepo.NewAuthRepository(r.orm)
	// newUserRepo := userRepo.NewUserRepository(r.orm)

	// newAuthService := authService.NewAuthService(newAuthRepo, newUserRepo)
	// newAuthMiddlewareService := middlewareService.NewAuthMiddlewareService(newAuthRepo)

	// newAuthHandler := authHandler.NewAuthHandler(newAuthService)

	// newAuthMiddleware := middleware.NewAuthMiddleware(newAuthMiddlewareService)

	authRoute := r.router.Group("/auth")
	{
		authRoute.POST("/register", r.h.AuthHandler.Register)
		authRoute.POST("/login", r.h.AuthHandler.Login)
		authRoute.POST("/logout", r.h.AuthHandler.Logout)
		authRoute.POST("/token/refresh", r.h.AuthHandler.CreateRefreshToken)
		// authRoute.POST("/logout", newAuthMiddleware.AuthGuard(), r.h.AuthHandler.Logout)
		// authRoute.POST("/token/refresh", newAuthMiddleware.AuthRefreshGuard(), r.h.AuthHandler.CreateRefreshToken)
	}
}
