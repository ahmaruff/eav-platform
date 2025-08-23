package user

import "net/http"

type Handler struct {
	service Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: *service}
}

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	// TODO: get user ID from session
	// TODO: get user data
	// TODO: render template
	w.Write([]byte("Dashboard - Coming Soon"))
}
