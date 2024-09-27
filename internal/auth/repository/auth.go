package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prajnasatryass/tic-be/internal/auth/domain"
	userDomain "github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/tokenutil"
)

type authRepository struct {
	db        *sqlx.DB
	tokenUtil tokenutil.TokenUtil
}

func NewAuthRepository(db *sqlx.DB) domain.AuthRepository {
	return &authRepository{
		db:        db,
		tokenUtil: tokenutil.NewTokenUtil(),
	}
}

func (ar *authRepository) CreateAccessToken(user *userDomain.User, secret string, ttl int) (string, error) {
	return ar.tokenUtil.CreateAccessToken(user, secret, ttl)
}

func (ar *authRepository) CreateRefreshToken(user *userDomain.User, secret string, ttl int) (string, error) {
	return ar.tokenUtil.CreateRefreshToken(user, secret, ttl)
}

func (ar *authRepository) StoreRefreshToken(record *domain.RefreshTokenRecord) error {
	_, err := ar.db.ExecContext(context.Background(), "INSERT INTO jwt_refresh_tokens VALUES($1, $2, $3)", record.Token, record.UserID, record.IgnoreAfter)
	return err
}

func (ar *authRepository) RetrieveRefreshToken(refreshToken string) (domain.RefreshTokenRecord, error) {
	var record domain.RefreshTokenRecord
	err := ar.db.GetContext(context.Background(), &record, "SELECT * FROM jwt_refresh_tokens WHERE token = $1 AND now() < ignore_after", refreshToken)
	return record, err
}

func (ar *authRepository) DeleteRefreshToken(refreshToken string) error {
	_, err := ar.db.ExecContext(context.Background(), "DELETE FROM jwt_refresh_tokens WHERE token = $1", refreshToken)
	return err
}

func (ar *authRepository) DeleteUserRefreshTokens(userID uuid.UUID) error {
	_, err := ar.db.ExecContext(context.Background(), "DELETE FROM jwt_refresh_tokens WHERE user_id = $1", userID.String())
	return err
}

func (ar *authRepository) DeleteExpiredRefreshTokens() error {
	_, err := ar.db.ExecContext(context.Background(), "DELETE FROM jwt_refresh_tokens WHERE now() > ignore_after")
	return err
}
