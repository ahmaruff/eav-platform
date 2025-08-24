package auth

import (
	"net/http"

	"github.com/ahmaruff/eav-platform/internal/user"
	"github.com/ahmaruff/eav-platform/templates"
)

type Handler struct {
	authService *Service
	userService *user.Service
}

func NewHandler(authService *Service, userService *user.Service) *Handler {
	return &Handler{
		authService: authService,
		userService: userService,
	}
}

func (h *Handler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	// render login form
	templates.LoginForm().Render(r.Context(), w)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	req := user.LoginRequest{
		Email:    email,
		Password: password,
	}

	user, err := h.userService.ValidateLogin(r.Context(), req)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	h.authService.CreateSession(w, r, user.ID)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *Handler) ShowRegister(w http.ResponseWriter, r *http.Request) {
	templates.RegisterForm().Render(r.Context(), w)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if email == "" || password == "" {
		http.Error(w, "Email and password required", http.StatusBadRequest)
		return
	}

	if password != confirmPassword {
		http.Error(w, "Password missmatch", http.StatusBadRequest)
		return
	}

	req := user.CreateUserRequest{
		Email:    email,
		Password: password,
	}

	user, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.authService.CreateSession(w, r, user.ID)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.authService.DestroySession(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
