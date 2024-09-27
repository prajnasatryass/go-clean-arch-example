package usecase

import (
	"database/sql"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/prajnasatryass/go-clean-arch-example/config"
	"github.com/prajnasatryass/go-clean-arch-example/internal/auth/domain"
	userDomain "github.com/prajnasatryass/go-clean-arch-example/internal/user/domain"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/apperror"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/hasher"
	"time"
)

var (
	errEmailNotLinked = func(email string) error {
		return errors.New(email + " is not linked to any user")
	}
	errGetUser                      = func(err error) error { return errors.New("get user error: " + err.Error()) }
	errPasswordIncorrect            = errors.New("password incorrect")
	errCreateAccessToken            = func(err error) error { return errors.New("create access token error: " + err.Error()) }
	errCreateRefreshToken           = func(err error) error { return errors.New("create refresh token error: " + err.Error()) }
	errStoreRefreshToken            = func(err error) error { return errors.New("store refresh token error: " + err.Error()) }
	errRefreshTokenInvalidOrExpired = errors.New("refresh token invalid or expired")
	errRefreshTokenNotLinked        = errors.New("refresh token not linked to any user")
	errDeleteRefreshToken           = func(err error) error { return errors.New("delete refresh token error: " + err.Error()) }
)

type authUsecase struct {
	authRepository domain.AuthRepository
	userRepository userDomain.UserRepository
	jwtCfg         config.JWTConfig
	hasher         hasher.Hasher
}

func NewAuthUsecase(authRepository domain.AuthRepository, userRepository userDomain.UserRepository, jwtCfg config.JWTConfig) domain.AuthUsecase {
	return &authUsecase{
		authRepository: authRepository,
		userRepository: userRepository,
		jwtCfg:         jwtCfg,
		hasher:         hasher.NewHasher(),
	}
}

func (au *authUsecase) Login(email, password string) (domain.LoginResponse, error) {
	user, err := au.userRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.LoginResponse{}, apperror.NotFound(errEmailNotLinked(email))
		}
		return domain.LoginResponse{}, apperror.InternalServerError(errGetUser(err))
	}

	if !au.hasher.MatchPassword(password, user.Password) {
		return domain.LoginResponse{}, apperror.Unauthorized(errPasswordIncorrect)
	}

	accessToken, err := au.authRepository.CreateAccessToken(&user, au.jwtCfg.AccessTokenSecret, au.jwtCfg.AccessTokenTTL)
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(errCreateAccessToken(err))
	}

	refreshToken, err := au.authRepository.CreateRefreshToken(&user, au.jwtCfg.RefreshTokenSecret, au.jwtCfg.RefreshTokenTTL)
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(errCreateRefreshToken(err))
	}

	err = au.authRepository.StoreRefreshToken(&domain.RefreshTokenRecord{
		Token:       refreshToken,
		UserID:      user.ID,
		IgnoreAfter: time.Now().Add(time.Duration(au.jwtCfg.RefreshTokenTTL) * time.Second),
	})
	if err != nil {
		return domain.LoginResponse{}, apperror.InternalServerError(errStoreRefreshToken(err))
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
		return domain.RefreshResponse{}, apperror.InternalServerError(errCreateAccessToken(err))
	}

	newRefreshToken, err := au.authRepository.CreateRefreshToken(&user, au.jwtCfg.RefreshTokenSecret, au.jwtCfg.RefreshTokenTTL)
	if err != nil {
		return domain.RefreshResponse{}, apperror.InternalServerError(errCreateRefreshToken(err))
	}

	err = au.authRepository.StoreRefreshToken(&domain.RefreshTokenRecord{
		Token:       newRefreshToken,
		UserID:      user.ID,
		IgnoreAfter: time.Now().Add(time.Duration(au.jwtCfg.RefreshTokenTTL) * time.Second),
	})
	if err != nil {
		return domain.RefreshResponse{}, apperror.InternalServerError(errStoreRefreshToken(err))
	}

	go func() {
		err = au.authRepository.DeleteRefreshToken(refreshToken)
		if err != nil {
			log.Error(errDeleteRefreshToken(err))
		}
	}()

	return domain.RefreshResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (au *authUsecase) Logout(refreshToken string) error {
	return au.authRepository.DeleteRefreshToken(refreshToken)
}
