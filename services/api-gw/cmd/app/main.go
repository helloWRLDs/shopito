package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"shopito/pkg/log"
	jsonutil "shopito/pkg/util/json"
	"shopito/services/api-gw/config"
	"shopito/services/api-gw/internal/delivery/middleware"
	usercontroller "shopito/services/api-gw/internal/delivery/user"
	userservice "shopito/services/api-gw/internal/service/users"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "shopito/services/api-gw/docs"
)

// @title Shopito API Gateway
// @version 1.0.0
// @description API Endpoints that connects GRPC and REST endpoints
// @host localhost:3000
// @BasePath /api/v1
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Init("api_gw")
	userService := userservice.New()
	// adminService := adminservice.New()
	// productService := productservice.New()
	// authService := authservice.New()

	userDelivery := usercontroller.New(userService)
	// adminDelivery := admincontroller.New(adminService)
	// productDelivery := productcontroller.New(productService)
	// authDelivery := authcontroller.New(authService)

	router := chi.NewRouter()

	router.Mount("/swagger", httpSwagger.WrapHandler)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		jsonutil.EncodeJson(w, 200, "pong")
	})

	router.With(middleware.LogRequest, middleware.SecureHeaders).
		Route("/api/v1", func(r chi.Router) {
			r.Mount("/users", userDelivery.Routes())
			// r.Mount("/admin", adminDelivery.Routes())
			// r.Mount("/products", productDelivery.ProductRoutes())
			// r.Mount("/auth", authDelivery.Routes())
		})

	srv := http.Server{
		Addr:         config.ADDR,
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		logrus.WithField("addr", config.ADDR).Info("server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-quit
	logrus.Println("Server is shutting down...")

	// Closing grpc connections
	userService.Close()
	// adminService.Close()
	// productService.Close()
	// authService.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}
	logrus.Info("Server exiting")
}
