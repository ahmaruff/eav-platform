package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmaruff/eav-platform/internal/auth"
	"github.com/ahmaruff/eav-platform/internal/infrastructure/repository"
	"github.com/ahmaruff/eav-platform/internal/shared"
	"github.com/ahmaruff/eav-platform/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	// Load config .env
	godotenv.Load()
	config := shared.LoadConfig()

	// Setup logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: config.GetLogLevel(),
	}))

	slog.SetDefault(logger)

	// Setup database
	db, err := shared.SetupDatabase(config.Database.Path)
	if err != nil {
		slog.Error("Failed to setup database", "error", err)
		os.Exit(1)
	}

	defer db.Close() // cleanup

	// Repositories
	userRepo := repository.NewUserSQLite(db)

	// Services
	userService := user.NewService(userRepo)
	authService := auth.NewService(db)

	// Handlers
	userHandler := user.NewHandler(userService)
	authHandler := auth.NewHandler(authService, userService)

	// Setup routes
	router := setupRoutes(authService, userHandler, authHandler)

	// Server setup
	srv := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Gracefull shutdown
	done := make(chan bool, 1)      // tanda kalau server sudah benar-benar mati
	quit := make(chan os.Signal, 1) // untuk menerima sinyal dari OS (Ctrl+C, SIGTERM)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Goroutine
	go func() {
		<-quit // blok sampai ada sinyal masuk ke channel quit
		slog.Info("Server shutting down")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Server forced to shutdown", "error", err)
		}
		close(done) // kasih tanda ke main goroutine kalau shutdown selesai

	}()

	slog.Info("Server starting", "addr", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}

	<-done
	slog.Info("Server stopped")
}
