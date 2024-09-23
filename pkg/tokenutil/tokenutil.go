package tokenutil

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/prajnasatryass/tic-be/internal/middleware"
	userDomain "github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/appconstants"
	"time"
)

func CreateAccessToken(user *userDomain.User, secret string, ttl int) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		Issuer:   appconstants.CompanySiteAddress,
		Subject:  user.ID.String(),
		Expiry:   now.Add(time.Duration(ttl) * time.Second),
		IssuedAt: now,
		Data: middleware.JWTClaimsData{
			ID:     user.ID,
			RoleID: user.RoleID,
		},
	})

	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(user *userDomain.User, secret string, ttl int) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.JWTClaims{
		Issuer:   appconstants.CompanySiteAddress,
		Subject:  user.ID.String(),
		Expiry:   now.Add(time.Duration(ttl) * time.Second),
		IssuedAt: now,
	})

	return token.SignedString([]byte(secret))
}
