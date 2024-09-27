package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/constants"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(email, password string) (uuid.UUID, error) {
	var newUserID uuid.UUID
	err := ur.db.GetContext(context.Background(), &newUserID, "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id", email, password)
	return newUserID, err
}

func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	err := ur.db.GetContext(context.Background(), &user, "SELECT * FROM users WHERE email = $1", email)
	return user, err
}

func (ur *userRepository) GetByID(id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := ur.db.GetContext(context.Background(), &user, "SELECT * FROM users WHERE id = $1", id)
	return user, err
}

func (ur *userRepository) UpdateRoleByID(id uuid.UUID, roleID constants.UserRole) error {
	_, err := ur.db.ExecContext(context.Background(), "UPDATE users SET role_id = $1, updated_at = now() WHERE id = $2", roleID, id)
	return err
}

func (ur *userRepository) DeleteByID(id uuid.UUID) error {
	_, err := ur.db.ExecContext(context.Background(), "UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL", id)
	return err
}

func (ur *userRepository) PermaDeleteByID(id uuid.UUID) error {
	_, err := ur.db.ExecContext(context.Background(), "DELETE FROM users WHERE id = $1", id)
	return err
}
