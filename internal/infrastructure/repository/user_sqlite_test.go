package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/ahmaruff/eav-platform/internal/infrastructure/repository"
	"github.com/ahmaruff/eav-platform/internal/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}

	schema := `
	CREATE TABLE users (
		id TEXT PRIMARY KEY,
		email TEXT NOT NULL,
		password_hash TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);`
	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	return db
}

func TestUserRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewUserSQLite(db)

	u := &user.User{
		ID:           "u1",
		Email:        "test@example.com",
		PasswordHash: "hashed_pw",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// pakai context
	ctx := context.Background()

	err := repo.Create(ctx, u)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// verify inserted
	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM users WHERE id = ?", u.ID)
	if err != nil {
		t.Fatalf("failed to query db: %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 user, got %d", count)
	}
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewUserSQLite(db)

	u := &user.User{
		ID:           "u1",
		Email:        "test@example.com",
		PasswordHash: "hashed_pw",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	ctx := context.Background()

	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got, err := repo.GetByEmail(ctx, "test@example.com")
	if err != nil {
		t.Fatalf("failed to query db: %v", err)
	}
	if got == nil {
		t.Fatalf("expected user, got nil")
	}
	if got.Email != "test@example.com" {
		t.Errorf("expected email %q, got %q", "test@example.com", got.Email)
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewUserSQLite(db)

	u := &user.User{
		ID:           "u1",
		Email:        "test@example.com",
		PasswordHash: "hashed_pw",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	ctx := context.Background()

	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got, err := repo.GetByID(ctx, "u1")
	if err != nil {
		t.Fatalf("failed to query db: %v", err)
	}
	if got == nil {
		t.Fatalf("expected user, got nil")
	}
	if got.ID != "u1" {
		t.Errorf("expected email %q, got %q", "u1", got.ID)
	}
}
