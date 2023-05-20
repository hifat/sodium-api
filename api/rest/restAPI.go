package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hifat/sodium-api/docs"
	"github.com/hifat/sodium-api/internal/database"
	"github.com/hifat/sodium-api/internal/routes"
	"github.com/hifat/sodium-api/internal/utils/validity"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func API() {
	/* --------------------------------- Init DB -------------------------------- */
	orm := database.PostgresDB()
	db, err := orm.DB()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = db.Ping()
		if err != nil {
			// The database connection is closed
			fmt.Println("The database connection is closed.")
		} else {
			// The database connection is still open
			fmt.Println("The database connection is open.")
		}
	}()
	defer db.Close()

	/* ---------------------------- Validator config ---------------------------- */

	validity.Register()

	/* ------------------------------- Swag config ------------------------------ */

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Sodium API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	/* --------------------------- Running API server --------------------------- */
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	r := routes.New(orm, api)
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
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}
