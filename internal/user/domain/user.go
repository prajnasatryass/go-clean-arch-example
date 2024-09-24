package domain

import (
	"database/sql"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Email     string       `json:"email"`
	Password  string       `json:"password,omitempty"`
	RoleID    int          `json:"roleID" db:"role_id"`
	IsActive  bool         `json:"isActive" db:"is_active"`
	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type CreateRequest struct {
	Email    string `form:"email" validate:"email,required"`
	Password string `form:"password" validate:"required"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type GetByIDResponse struct {
	User `json:",inline"`
}

type DeleteByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type UserUsecase interface {
	Create(email, password string) (CreateResponse, error)
	GetByID(id uuid.UUID) (GetByIDResponse, error)
	DeleteByID(id uuid.UUID) error
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID(id uuid.UUID) (User, error)
	DeleteByID(id uuid.UUID) error
	PermaDeleteByID(id uuid.UUID) error
}
