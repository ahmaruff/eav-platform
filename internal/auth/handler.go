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
	var errors []string
	templates.LoginForm(errors).Render(r.Context(), w)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	var errors []string

	if email == "" {
		errors = append(errors, "Email is required")
	}

	if password == "" {
		errors = append(errors, "Password is required")
	}

	if len(errors) > 0 {
		templates.LoginForm(errors).Render(r.Context(), w)
		return
	}

	req := user.LoginRequest{
		Email:    email,
		Password: password,
	}

	user, err := h.userService.ValidateLogin(r.Context(), req)
	if err != nil {
		errors = append(errors, "Invalid email or password")
		templates.LoginForm(errors).Render(r.Context(), w)
		return
	}

	h.authService.CreateSession(w, r, user.ID)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *Handler) ShowRegister(w http.ResponseWriter, r *http.Request) {
	var errors []string
	templates.RegisterForm(errors).Render(r.Context(), w)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	var errors []string

	if email == "" {
		errors = append(errors, "Email is required")
	}

	if password == "" {
		errors = append(errors, "Password is required")
	}

	if confirmPassword == "" {
		errors = append(errors, "Confirm Password is required")
	}

	if len(errors) > 0 {
		templates.RegisterForm(errors).Render(r.Context(), w)
		return
	}

	if password != confirmPassword {
		errors = append(errors, "Password missmatch")
	}

	if len(errors) > 0 {
		templates.RegisterForm(errors).Render(r.Context(), w)
		return
	}

	req := user.CreateUserRequest{
		Email:    email,
		Password: password,
	}

	user, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		errors = append(errors, err.Error())
		templates.RegisterForm(errors).Render(r.Context(), w)
		return
	}

	h.authService.CreateSession(w, r, user.ID)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.authService.DestroySession(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
