package domain

import (
	"database/sql"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	RoleID    int          `db:"role_id"`
	IsActive  bool         `db:"is_active"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type UserUsecase interface {
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID(id uuid.UUID) (User, error)
	DeleteByID(id uuid.UUID) error
	PermaDeleteByID(id uuid.UUID) error
}
