package usecase

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/prajnasatryass/go-clean-arch-example/config"
	"github.com/prajnasatryass/go-clean-arch-example/internal/auth/domain"
	userDomain "github.com/prajnasatryass/go-clean-arch-example/internal/user/domain"
	mockDomain "github.com/prajnasatryass/go-clean-arch-example/mocks/internal_/auth/domain"
	mockUserDomain "github.com/prajnasatryass/go-clean-arch-example/mocks/internal_/user/domain"
	mockHasher "github.com/prajnasatryass/go-clean-arch-example/mocks/pkg/hasher"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/hasher"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

var (
	jwtCfg = config.JWTConfig{
		AccessTokenSecret:  "access_token_secret",
		AccessTokenTTL:     1,
		RefreshTokenSecret: "refresh_token_secret",
		RefreshTokenTTL:    2,
	}

	inputUser = userDomain.User{
		Email:    "user@example.com",
		Password: "123",
	}
	userID    = uuid.New()
	matchUser = userDomain.User{
		ID:       userID,
		Email:    "user@example.com",
		Password: "$2a$10$f8Ysjht6jVtmpCZHKONOqemdBaSfJdxKEtrBnwnSWprWBYhx3Kiee", // 123
	}

	accessToken  = "access_token"
	refreshToken = "refresh_token"

	refreshTokenRecord = domain.RefreshTokenRecord{
		Token:  refreshToken,
		UserID: userID,
	}
)

func TestNewAuthUsecase(t *testing.T) {
	type args struct {
		authRepository domain.AuthRepository
		userRepository userDomain.UserRepository
		jwtCfg         config.JWTConfig
		hasher         hasher.Hasher
	}
	tests := []struct {
		name string
		args args
		want domain.AuthUsecase
	}{
		{
			name: "success",
			args: args{
				authRepository: nil,
				userRepository: nil,
				jwtCfg:         config.JWTConfig{},
			},
			want: &authUsecase{
				authRepository: nil,
				userRepository: nil,
				jwtCfg:         config.JWTConfig{},
				hasher:         hasher.NewHasher(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecase(tt.args.authRepository, tt.args.userRepository, tt.args.jwtCfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUsecase_Login(t *testing.T) {
	mockAuthRepository := mockDomain.NewMockAuthRepository(t)
	mockUserRepository := mockUserDomain.NewMockUserRepository(t)
	mockTestHasher := mockHasher.NewMockHasher(t)

	type fields struct {
		authRepository domain.AuthRepository
		userRepository userDomain.UserRepository
		jwtCfg         config.JWTConfig
		hasher         hasher.Hasher
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    domain.LoginResponse
		wantErr bool
	}{
		{
			name: "get user by email no result",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(userDomain.User{}, sql.ErrNoRows).Once()
			},
			want:    domain.LoginResponse{},
			wantErr: true,
		},
		{
			name: "get user by email error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(userDomain.User{}, errors.New("")).Once()
			},
			want:    domain.LoginResponse{},
			wantErr: true,
		},
		{
			name: "password does not match",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(matchUser, nil).Once()
				mockTestHasher.EXPECT().MatchPassword(inputUser.Password, matchUser.Password).Return(false).Once()
			},
			want:    domain.LoginResponse{},
			wantErr: true,
		},
		{
			name: "create access token error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(matchUser, nil).Once()
				mockTestHasher.EXPECT().MatchPassword(inputUser.Password, matchUser.Password).Return(true).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return("", errors.New("")).Once()
			},
			want:    domain.LoginResponse{},
			wantErr: true,
		},
		{
			name: "create refresh token error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(matchUser, nil).Once()
				mockTestHasher.EXPECT().MatchPassword(inputUser.Password, matchUser.Password).Return(true).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return(accessToken, nil).Once()
				mockAuthRepository.EXPECT().CreateRefreshToken(&matchUser, jwtCfg.RefreshTokenSecret, jwtCfg.RefreshTokenTTL).Return("", errors.New("")).Once()
			},
			want:    domain.LoginResponse{},
			wantErr: true,
		},
		{
			name: "store refresh token error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(matchUser, nil).Once()
				mockTestHasher.EXPECT().MatchPassword(inputUser.Password, matchUser.Password).Return(true).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return(accessToken, nil).Once()
				mockAuthRepository.EXPECT().CreateRefreshToken(&matchUser, jwtCfg.RefreshTokenSecret, jwtCfg.RefreshTokenTTL).Return(refreshToken, nil).Once()
				mockAuthRepository.EXPECT().StoreRefreshToken(mock.Anything).Return(errors.New("")).Once()
			},
			want:    domain.LoginResponse{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(matchUser, nil).Once()
				mockTestHasher.EXPECT().MatchPassword(inputUser.Password, matchUser.Password).Return(true).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return(accessToken, nil).Once()
				mockAuthRepository.EXPECT().CreateRefreshToken(&matchUser, jwtCfg.RefreshTokenSecret, jwtCfg.RefreshTokenTTL).Return(refreshToken, nil).Once()
				mockAuthRepository.EXPECT().StoreRefreshToken(mock.Anything).Return(nil).Once()
			},
			want: domain.LoginResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &authUsecase{
				authRepository: tt.fields.authRepository,
				userRepository: tt.fields.userRepository,
				jwtCfg:         tt.fields.jwtCfg,
				hasher:         tt.fields.hasher,
			}
			tt.pre()
			got, err := au.Login(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUsecase_Logout(t *testing.T) {
	mockAuthRepository := mockDomain.NewMockAuthRepository(t)
	mockUserRepository := mockUserDomain.NewMockUserRepository(t)

	type fields struct {
		authRepository domain.AuthRepository
		userRepository userDomain.UserRepository
		jwtCfg         config.JWTConfig
		hasher         hasher.Hasher
	}
	type args struct {
		refreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().DeleteRefreshToken(refreshToken).Return(nil).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &authUsecase{
				authRepository: tt.fields.authRepository,
				userRepository: tt.fields.userRepository,
				jwtCfg:         tt.fields.jwtCfg,
			}
			tt.pre()
			if err := au.Logout(tt.args.refreshToken); (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_authUsecase_Refresh(t *testing.T) {
	mockAuthRepository := mockDomain.NewMockAuthRepository(t)
	mockUserRepository := mockUserDomain.NewMockUserRepository(t)

	type fields struct {
		authRepository domain.AuthRepository
		userRepository userDomain.UserRepository
		jwtCfg         config.JWTConfig
		hasher         hasher.Hasher
	}
	type args struct {
		refreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    domain.RefreshResponse
		wantErr bool
	}{
		{
			name: "retrieve refresh token no result",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(domain.RefreshTokenRecord{}, sql.ErrNoRows).Once()
			},
			want:    domain.RefreshResponse{},
			wantErr: true,
		},
		{
			name: "retrieve refresh token error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(domain.RefreshTokenRecord{}, errors.New("")).Once()
			},
			want:    domain.RefreshResponse{},
			wantErr: true,
		},
		{
			name: "get user by ID no result",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(refreshTokenRecord, nil).Once()
				mockUserRepository.EXPECT().GetByID(refreshTokenRecord.UserID).Return(userDomain.User{}, sql.ErrNoRows).Once()
			},
			want:    domain.RefreshResponse{},
			wantErr: true,
		},
		{
			name: "get user by ID error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(refreshTokenRecord, nil).Once()
				mockUserRepository.EXPECT().GetByID(refreshTokenRecord.UserID).Return(userDomain.User{}, errors.New("")).Once()
			},
			want:    domain.RefreshResponse{},
			wantErr: true,
		},
		{
			name: "create access token error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(refreshTokenRecord, nil).Once()
				mockUserRepository.EXPECT().GetByID(refreshTokenRecord.UserID).Return(matchUser, nil).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return("", errors.New("")).Once()
			},
			want:    domain.RefreshResponse{},
			wantErr: true,
		},
		{
			name: "create refresh token error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(refreshTokenRecord, nil).Once()
				mockUserRepository.EXPECT().GetByID(refreshTokenRecord.UserID).Return(matchUser, nil).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return(accessToken, nil).Once()
				mockAuthRepository.EXPECT().CreateRefreshToken(&matchUser, jwtCfg.RefreshTokenSecret, jwtCfg.RefreshTokenTTL).Return("", errors.New("")).Once()
			},
			want:    domain.RefreshResponse{},
			wantErr: true,
		},
		{
			name: "store refresh token error",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(refreshTokenRecord, nil).Once()
				mockUserRepository.EXPECT().GetByID(refreshTokenRecord.UserID).Return(matchUser, nil).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return(accessToken, nil).Once()
				mockAuthRepository.EXPECT().CreateRefreshToken(&matchUser, jwtCfg.RefreshTokenSecret, jwtCfg.RefreshTokenTTL).Return(refreshToken, nil).Once()
				mockAuthRepository.EXPECT().StoreRefreshToken(mock.Anything).Return(errors.New("")).Once()
			},
			want:    domain.RefreshResponse{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				authRepository: mockAuthRepository,
				userRepository: mockUserRepository,
				jwtCfg:         jwtCfg,
			},
			args: args{
				refreshToken: refreshToken,
			},
			pre: func() {
				mockAuthRepository.EXPECT().RetrieveRefreshToken(refreshToken).Return(refreshTokenRecord, nil).Once()
				mockUserRepository.EXPECT().GetByID(refreshTokenRecord.UserID).Return(matchUser, nil).Once()
				mockAuthRepository.EXPECT().CreateAccessToken(&matchUser, jwtCfg.AccessTokenSecret, jwtCfg.AccessTokenTTL).Return(accessToken, nil).Once()
				mockAuthRepository.EXPECT().CreateRefreshToken(&matchUser, jwtCfg.RefreshTokenSecret, jwtCfg.RefreshTokenTTL).Return(refreshToken, nil).Once()
				mockAuthRepository.EXPECT().StoreRefreshToken(mock.Anything).Return(nil).Once()
				mockAuthRepository.EXPECT().DeleteRefreshToken(refreshToken).Return(nil).Once()
			},
			want: domain.RefreshResponse{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			au := &authUsecase{
				authRepository: tt.fields.authRepository,
				userRepository: tt.fields.userRepository,
				jwtCfg:         tt.fields.jwtCfg,
			}
			tt.pre()
			got, err := au.Refresh(tt.args.refreshToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Refresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Refresh() got = %v, want %v", got, tt.want)
			}
		})
	}
}
