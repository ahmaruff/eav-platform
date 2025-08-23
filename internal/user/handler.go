package user

import (
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: *service}
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	// Safe type assertion
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Use service, not repo directly
	user, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Fix string formatting
	w.Write([]byte("Dashboard - " + user.Email))
}
