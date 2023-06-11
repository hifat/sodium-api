package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/docs"
	"github.com/hifat/sodium-api/internal/di"
	"github.com/hifat/sodium-api/internal/routes"
	"github.com/hifat/sodium-api/internal/utils/validity"
	_ "github.com/joho/godotenv/autoload"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	if os.Getenv("APP_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	/* --------------------------------- Init DB -------------------------------- */
	handlerWire, cleanup := di.InitializeAPI()
	defer cleanup()

	/* ---------------------------- Validator config ---------------------------- */
	validity.Register()

	/* ------------------------------- Swag config ------------------------------ */

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Sodium API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	/* --------------------------- Running API server --------------------------- */
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
	}

	router := gin.Default()

	router.Use(cors.New(corsConfig))
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	r := routes.New(api, handlerWire)
	r.Register()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:           os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Println(err)
	}
}
