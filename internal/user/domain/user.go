package domain

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/prajnasatryass/tic-be/pkg/constants"
)

type User struct {
	ID        uuid.UUID          `json:"id"`
	Email     string             `json:"email"`
	Password  string             `json:"password,omitempty"`
	RoleID    constants.UserRole `json:"roleID" db:"role_id"`
	CreatedAt sql.NullTime       `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime       `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime       `json:"deleted_at" db:"deleted_at"`
}

type CreateRequest struct {
	Email    string `form:"email" validate:"email,required"`
	Password string `form:"password" validate:"required"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

type ByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type UpdateRoleByIDRequest struct {
	ID     uuid.UUID          `param:"id" validate:"required"`
	RoleID constants.UserRole `json:"roleID" validate:"required"`
}

type GetByIDResponse struct {
	User `json:",inline"`
}

type UserUsecase interface {
	Create(email, password string) (CreateResponse, error)
	GetByID(id uuid.UUID) (GetByIDResponse, error)
	UpdateRoleByID(id uuid.UUID, roleID constants.UserRole) error
	DeleteByID(id uuid.UUID) error
}

type UserRepository interface {
	Create(email, password string) (uuid.UUID, error)
	GetByEmail(email string) (User, error)
	GetByID(id uuid.UUID) (User, error)
	UpdateRoleByID(id uuid.UUID, roleID constants.UserRole) error
	DeleteByID(id uuid.UUID) error
	PermaDeleteByID(id uuid.UUID) error
}
