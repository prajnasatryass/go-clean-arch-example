package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"tic-be/internal/user/domain"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *domain.User) error {
	_, err := ur.db.ExecContext(context.Background(), "INSERT INTO users (email, password, role_id) VALUES ($1, $2, $3)", user.Email, user.Password, user.RoleID)
	return err
}

func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := ur.db.GetContext(context.Background(), &user, "SELECT * FROM users WHERE deleted_at IS NULL AND email = $1", email)
	return user, err
}

func (ur *userRepository) GetByID(id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := ur.db.GetContext(context.Background(), &user, "SELECT * FROM users WHERE deleted_at IS NULL AND id = $1", id)
	return user, err
}

func (ur *userRepository) DeleteByID(id uuid.UUID) error {
	_, err := ur.db.ExecContext(context.Background(), "UPDATE users SET deleted_at = NOW() WHERE id = $1", id)
	return err
}

func (ur *userRepository) PermaDeleteByID(id uuid.UUID) error {
	_, err := ur.db.ExecContext(context.Background(), "DELETE FROM users WHERE id = $1", id)
	return err
}
