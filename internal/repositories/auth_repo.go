package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"restapp/internal/models"
)

type AuthRepositoryInterface interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Register(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := r.db.ExecContext(ctx,
		"INSERT INTO users (username, password, email, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		user.Username,
		user.Password,
		user.Email,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register user, error: %s", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve inserted user ID: %w", err)
	}

	user.Id = int(id)

	return user, nil
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.GetContext(ctx,
		&user,
		"SELECT id, username, password, email, role, created_at, updated_at FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("failed to get user by email, error: %s", err)
	}

	return &user, nil
}
