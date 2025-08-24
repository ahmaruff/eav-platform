package main

import (
	"github.com/ahmaruff/eav-platform/internal/auth"
	"github.com/ahmaruff/eav-platform/internal/user"
	"github.com/ahmaruff/eav-platform/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
)

func customRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Panic recovered", "error", err)
				w.WriteHeader(http.StatusInternalServerError)
				templates.Error500().Render(r.Context(), w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func setupRoutes(authService *auth.Service, userHandler *user.Handler, authHandler *auth.Handler) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(authService.SessionMiddleware)
	r.Use(customRecoverer) // Custom 500 handler

	// STATIC
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Error routes
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		templates.Error404().Render(r.Context(), w)
	})

	// Basic routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Use(authService.RedirectIfAuthenticated)
		r.Get("/login", authHandler.ShowLogin)
		r.Post("/login", authHandler.Login)
		r.Get("/register", authHandler.ShowRegister)
		r.Post("/register", authHandler.Register)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(authService.RequireAuth)
		r.Get("/dashboard", userHandler.Dashboard)
		r.Post("/logout", authHandler.Logout)
	})

	return r
}
