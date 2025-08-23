package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ahmaruff/eav-platform/internal/user"
	"github.com/jmoiron/sqlx"
)

var ErrUserNotFound = errors.New("User not found")

type UserSQLite struct {
	db *sqlx.DB
}

func NewUserSQLite(db *sqlx.DB) user.Repository {
	return &UserSQLite{db: db}
}

func (r *UserSQLite) Create(ctx context.Context, u *user.User) error {
	now := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now

	query := `
        	INSERT INTO users(id, email, password_hash, created_at, updated_at)
        	VALUES (:id, :email, :password_hash, :created_at, :updated_at)
    	`
	_, err := r.db.NamedExecContext(ctx, query, u)
	if err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}
	return nil
}

func (r *UserSQLite) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	query := `
		SELECT id, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = ?
	`

	err := r.db.GetContext(ctx, &u, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("Failed to get user by email: %w", err)
	}

	return &u, nil
}

func (r *UserSQLite) GetByID(ctx context.Context, id string) (*user.User, error) {
	var u user.User
	query := `
		SELECT id, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	err := r.db.GetContext(ctx, &u, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("Failed to get user by id: %w", err)
	}

	return &u, nil
}
