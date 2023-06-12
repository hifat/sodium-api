package routes

func (r routes) authRoute() {
	authRoute := r.router.Group("/auth")
	{
		authRoute.POST("/register", r.handler.AuthHandler.Register)
		authRoute.POST("/login", r.handler.AuthHandler.Login)
		authRoute.POST("/logout", r.middleware.AuthGuard(), r.handler.AuthHandler.Logout)
		authRoute.POST("/token/refresh", r.middleware.AuthRefreshGuard(), r.handler.AuthHandler.CreateRefreshToken)
	}
}
