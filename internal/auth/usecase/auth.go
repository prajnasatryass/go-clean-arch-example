package usecase

import (
	"database/sql"
	"errors"
	"github.com/prajnasatryass/tic-be/config"
	"github.com/prajnasatryass/tic-be/internal/auth/domain"
	userDomain "github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/apperror"
	"github.com/prajnasatryass/tic-be/pkg/hasher"
	"time"
)

var (
	errEmailNotLinked = func(email string) error {
		return errors.New(email + " is not linked to any user")
	}
	errPasswordIncorrect            = errors.New("password incorrect")
	errRefreshTokenInvalidOrExpired = errors.New("refresh token invalid or expired")
	errRefreshTokenNotLinked        = errors.New("refresh token not linked to any user")
)

type authUsecase struct {
	authRepository domain.AuthRepository
	userRepository userDomain.UserRepository
	jwtCfg         config.JWTConfig
}

func NewAuthUsecase(authRepository domain.AuthRepository, userRepository userDomain.UserRepository, jwtCfg config.JWTConfig) domain.AuthUsecase {
	return &authUsecase{
		authRepository: authRepository,
		userRepository: userRepository,
		jwtCfg:         jwtCfg,
	}
}

func (au *authUsecase) Login(email, password string) (domain.LoginResponse, error) {
	user, err := au.userRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.LoginResponse{}, apperror.NotFound(errEmailNotLinked(email))
		}
		return domain.LoginResponse{}, apperror.InternalServerError(err)
	}

	if !hasher.MatchPassword(password, user.Password) {
		return domain.LoginResponse{}, apperror.Unauthorized(errPasswordIncorrect)
	}

	accessToken, err := au.authRepository.CreateAccessToken(&user, au.jwtCfg.AccessTokenSecret, au.jwtCfg.AccessTokenTTL)
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(err)
	}

	refreshToken, err := au.authRepository.CreateRefreshToken(&user, au.jwtCfg.RefreshTokenSecret, au.jwtCfg.RefreshTokenTTL)
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(err)
	}

	err = au.authRepository.StoreRefreshToken(&domain.RefreshTokenRecord{
		Token:       refreshToken,
		UserID:      user.ID,
		IgnoreAfter: time.Now().Add(time.Duration(au.jwtCfg.RefreshTokenTTL) * time.Second),
	})
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(err)
	}

	return domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (au *authUsecase) Refresh(refreshToken string) (domain.RefreshResponse, error) {
	record, err := au.authRepository.RetrieveRefreshToken(refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.RefreshResponse{}, apperror.NotFound(errRefreshTokenInvalidOrExpired)
		}
		return domain.RefreshResponse{}, err
	}

	user, err := au.userRepository.GetByID(record.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.RefreshResponse{}, apperror.NotFound(errRefreshTokenNotLinked)
		}
		return domain.RefreshResponse{}, err
	}

	newAccessToken, err := au.authRepository.CreateAccessToken(&user, au.jwtCfg.AccessTokenSecret, au.jwtCfg.AccessTokenTTL)
	if err != nil {
		return domain.RefreshResponse{}, apperror.InternalServerError(err)
	}

	newRefreshToken, err := au.authRepository.CreateRefreshToken(&user, au.jwtCfg.RefreshTokenSecret, au.jwtCfg.RefreshTokenTTL)
	if err != nil {
		return domain.RefreshResponse{}, apperror.InternalServerError(err)
	}

	err = au.authRepository.StoreRefreshToken(&domain.RefreshTokenRecord{
		Token:       newRefreshToken,
		UserID:      user.ID,
		IgnoreAfter: time.Now().Add(time.Duration(au.jwtCfg.RefreshTokenTTL) * time.Second),
	})
	if err != nil {
		return domain.RefreshResponse{}, apperror.InternalServerError(err)
	}

	return domain.RefreshResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (au *authUsecase) Logout(refreshToken string) error {
	return au.authRepository.DeleteRefreshToken(refreshToken)
}
