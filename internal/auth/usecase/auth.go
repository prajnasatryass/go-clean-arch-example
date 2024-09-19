package usecase

import (
	"database/sql"
	"errors"
	"tic-be/config"
	"tic-be/internal/auth/domain"
	userDomain "tic-be/internal/user/domain"
	"tic-be/pkg/apperror"
	"tic-be/pkg/hasher"
	"tic-be/pkg/tokenutil"
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
	cfg            config.Config
}

func NewAuthUsecase(authRepository domain.AuthRepository, userRepository userDomain.UserRepository, cfg config.Config) domain.AuthUsecase {
	return &authUsecase{
		authRepository: authRepository,
		userRepository: userRepository,
		cfg:            cfg,
	}
}

func (au *authUsecase) Login(email, password string) (domain.LoginResponse, error) {
	user, err := au.userRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.LoginResponse{}, apperror.Unauthorized(errEmailNotLinked(email))
		}
		return domain.LoginResponse{}, apperror.InternalServerError(err)
	}

	if !hasher.MatchPassword(password, user.Password) {
		return domain.LoginResponse{}, apperror.Unauthorized(errPasswordIncorrect)
	}

	accessToken, err := tokenutil.CreateAccessToken(&user, au.cfg.JWT.AccessTokenSecret, au.cfg.JWT.AccessTokenTTL)
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(err)
	}

	refreshToken, err := tokenutil.CreateRefreshToken(&user, au.cfg.JWT.RefreshTokenSecret, au.cfg.JWT.RefreshTokenTTL)
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(err)
	}

	err = au.authRepository.StoreRefreshToken(&domain.RefreshTokenRecord{
		Token:       refreshToken,
		UserID:      user.ID,
		IgnoreAfter: time.Now().Add(time.Duration(au.cfg.JWT.RefreshTokenTTL) * time.Second),
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

	newAccessToken, err := tokenutil.CreateAccessToken(&user, au.cfg.JWT.AccessTokenSecret, au.cfg.JWT.AccessTokenTTL)
	if err != nil {
		return domain.RefreshResponse{}, apperror.InternalServerError(err)
	}

	newRefreshToken, err := tokenutil.CreateRefreshToken(&user, au.cfg.JWT.RefreshTokenSecret, au.cfg.JWT.RefreshTokenTTL)
	if err != nil {
		return domain.RefreshResponse{}, apperror.InternalServerError(err)
	}

	err = au.authRepository.StoreRefreshToken(&domain.RefreshTokenRecord{
		Token:       newRefreshToken,
		UserID:      user.ID,
		IgnoreAfter: time.Now().Add(time.Duration(au.cfg.JWT.RefreshTokenTTL) * time.Second),
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
