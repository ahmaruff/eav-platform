package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ahmaruff/eav-platform/internal/shared"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load config .env
	godotenv.Load()
	config := shared.Load()

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

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Basic routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("EAV Platform - Coming soon!"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Server setup
	srv := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      r,
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
