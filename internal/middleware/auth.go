package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/prajnasatryass/tic-be/config"
	"github.com/prajnasatryass/tic-be/pkg/apperror"
	"github.com/prajnasatryass/tic-be/pkg/appresponse"
	"github.com/prajnasatryass/tic-be/pkg/constants"
	"time"
)

type JWTClaims struct {
	Issuer    string        `json:"iss"`
	Subject   string        `json:"sub"`
	Audience  []string      `json:"aud"`
	Expiry    time.Time     `json:"exp"`
	NotBefore time.Time     `json:"nbf"`
	IssuedAt  time.Time     `json:"iat"`
	Data      JWTClaimsData `json:"data"`
}
type JWTClaimsData struct {
	Email  string             `json:"email"`
	RoleID constants.UserRole `json:"roleID"`
}

func (j *JWTClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(j.Expiry), nil
}
func (j *JWTClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(j.IssuedAt), nil
}
func (j *JWTClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(j.NotBefore), nil
}
func (j *JWTClaims) GetIssuer() (string, error) {
	return j.Issuer, nil
}
func (j *JWTClaims) GetSubject() (string, error) {
	return j.Subject, nil
}
func (j *JWTClaims) GetAudience() (jwt.ClaimStrings, error) {
	return j.Audience, nil
}

const (
	authorizationHeaderKey    = "Authorization"
	authorizationHeaderPrefix = "Bearer "
	jwtClaimsKey              = "jwtClaims"
)

var (
	errAuthorizationHeaderMissingOrEmpty = errors.New("authorization header missing or empty")
	errJWTTokenSigningMethodUnexpected   = func(signingMethod interface{}) error {
		return fmt.Errorf("JWT token signing method unexpected: %v", signingMethod)
	}
	errJWTTokenInvalid = func(errStr string) error {
		return errors.New("JWT token invalid: " + errStr)
	}
	errRoleUnassigned        = errors.New("user role unassigned. Access rejected")
	errJWTClaimsDataNotFound = errors.New("JWT claims data not found. Please login again")
)

func JWTAuth(cfg config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(authorizationHeaderKey)
			if authHeader == "" {
				return appresponse.ErrorResponseBuilder(apperror.Forbidden(errAuthorizationHeaderMissingOrEmpty)).Return(c)
			}
			tokenStr := authHeader[len(authorizationHeaderPrefix):]

			token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errJWTTokenSigningMethodUnexpected(token.Header["alg"])
				}
				return []byte(cfg.JWT.AccessTokenSecret), nil
			})
			if err != nil {
				return appresponse.ErrorResponseBuilder(apperror.Forbidden(errJWTTokenInvalid(err.Error()))).Return(c)
			}

			claims := token.Claims.(*JWTClaims)
			if claims.Data.RoleID == constants.UserRoleUnassigned {
				return appresponse.ErrorResponseBuilder(apperror.Forbidden(errRoleUnassigned)).Return(c)
			}
			c.Set(jwtClaimsKey, claims)

			return next(c)
		}
	}
}

func GetJWTClaimsData(c echo.Context) (*JWTClaimsData, error) {
	data, ok := c.Get(jwtClaimsKey).(*JWTClaimsData)
	if !ok {
		return nil, apperror.Unauthorized(errJWTClaimsDataNotFound)
	}

	return data, nil
}
