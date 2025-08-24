package main

import (
	"github.com/ahmaruff/eav-platform/internal/auth"
	"github.com/ahmaruff/eav-platform/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func setupRoutes(authService *auth.Service, userHandler *user.Handler, authHandler *auth.Handler) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(authService.SessionMiddleware)

	// STATIC
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

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
