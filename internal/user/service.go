package user

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service struct {
	Repo     Repository
	validate *validator.Validate
}

func NewService(repo Repository) *Service {
	validate := validator.New()
	return &Service{Repo: repo, validate: validate}
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("Validation failed: %w", err)
	}

	user, _ := s.Repo.GetByEmail(ctx, req.Email)
	if user != nil {
		return nil, fmt.Errorf("Email already taken")
	}

	id := uuid.New()
	var u User
	now := time.Now()

	u.ID = id.String()
	u.Email = req.Email
	u.CreatedAt = now
	u.UpdatedAt = now
	err := u.SetPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("Failed to set password %w", err)
	}

	err = s.Repo.Create(ctx, &u)
	if err != nil {
		return nil, fmt.Errorf("Failed to create user %w", err)
	}

	return &u, nil
}

func (s *Service) ValidateLogin(ctx context.Context, req LoginRequest) (*User, error) {
	if err := s.validate.Struct(req); err != nil {
		return nil, fmt.Errorf("Validation failed: %w", err)
	}

	user, err := s.Repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	err = user.CheckPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*User, error) {
	user, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("User not found: %w", err)
	}

	return user, nil
}
