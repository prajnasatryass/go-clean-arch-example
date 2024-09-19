package domain

import (
	"github.com/google/uuid"
	"time"
)

type RefreshTokenRecord struct {
	Token       string
	UserID      uuid.UUID `db:"user_id"`
	IgnoreAfter time.Time `db:"ignore_after"`
}

type LoginRequest struct {
	Email    string `form:"email" validate:"email,required"`
	Password string `form:"password" validate:"required"`
}
type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
type RefreshResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type AuthUsecase interface {
	Login(email, password string) (LoginResponse, error)
	Refresh(refreshToken string) (RefreshResponse, error)
	Logout(refreshToken string) error
}

type AuthRepository interface {
	StoreRefreshToken(record *RefreshTokenRecord) error
	RetrieveRefreshToken(refreshToken string) (RefreshTokenRecord, error)
	DeleteRefreshToken(refreshToken string) error
	DeleteUserRefreshTokens(userID uuid.UUID) error
	DeleteExpiredRefreshTokens() error
}
