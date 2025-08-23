package auth

import (
	"context"
	"net/http"
)

func (s *Service) SessionMiddleware(next http.Handler) http.Handler {
	return s.sessionManager.LoadAndSave(next)
}

func (s *Service) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := s.GetUserID(r)
		if userID == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// Add userID to context
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Service) RedirectIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := s.GetUserID(r)
		if userID != "" { // if authenticated
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r) // continue to login/register
	})
}
