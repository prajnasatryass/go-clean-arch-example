package usecase

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/prajnasatryass/tic-be/internal/user/domain"
	mockUserDomain "github.com/prajnasatryass/tic-be/mocks/internal_/user/domain"
	mockHasher "github.com/prajnasatryass/tic-be/mocks/pkg/hasher"
	"github.com/prajnasatryass/tic-be/pkg/constants"
	"github.com/prajnasatryass/tic-be/pkg/hasher"
	"reflect"
	"testing"
)

var (
	inputUser = domain.User{
		Email:    "user@ticindo.com",
		Password: "123",
	}
	hashedPassword = "$2a$10$f8Ysjht6jVtmpCZHKONOqemdBaSfJdxKEtrBnwnSWprWBYhx3Kiee" // 123
	newUserID      = uuid.New()
	newRoleID      = constants.UserRoleRoot
	badRoleID      = 9999
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		userRepository domain.UserRepository
	}
	tests := []struct {
		name string
		args args
		want domain.UserUsecase
	}{
		{
			name: "success",
			args: args{
				userRepository: nil,
			},
			want: &userUsecase{
				userRepository: nil,
				hasher:         hasher.NewHasher(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_Create(t *testing.T) {
	mockUserRepository := mockUserDomain.NewMockUserRepository(t)
	mockTestHasher := mockHasher.NewMockHasher(t)

	type fields struct {
		userRepository domain.UserRepository
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
		want    domain.CreateResponse
		wantErr bool
	}{
		{
			name: "get user by id error",
			fields: fields{
				userRepository: mockUserRepository,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(domain.User{}, errors.New("")).Once()
			},
			want:    domain.CreateResponse{},
			wantErr: true,
		},
		{
			name: "get user by id found",
			fields: fields{
				userRepository: mockUserRepository,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(domain.User{Email: inputUser.Email}, nil).Once()
			},
			want:    domain.CreateResponse{},
			wantErr: true,
		},
		{
			name: "hash password error",
			fields: fields{
				userRepository: mockUserRepository,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(domain.User{}, sql.ErrNoRows).Once()
				mockTestHasher.EXPECT().HashPassword(inputUser.Password).Return("", errors.New("")).Once()
			},
			want:    domain.CreateResponse{},
			wantErr: true,
		},
		{
			name: "create user error",
			fields: fields{
				userRepository: mockUserRepository,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(domain.User{}, sql.ErrNoRows).Once()
				mockTestHasher.EXPECT().HashPassword(inputUser.Password).Return(hashedPassword, nil).Once()
				mockUserRepository.EXPECT().Create(inputUser.Email, hashedPassword).Return(uuid.UUID{}, errors.New("")).Once()
			},
			want:    domain.CreateResponse{},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				userRepository: mockUserRepository,
				hasher:         mockTestHasher,
			},
			args: args{
				email:    inputUser.Email,
				password: inputUser.Password,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByEmail(inputUser.Email).Return(domain.User{}, sql.ErrNoRows).Once()
				mockTestHasher.EXPECT().HashPassword(inputUser.Password).Return(hashedPassword, nil).Once()
				mockUserRepository.EXPECT().Create(inputUser.Email, hashedPassword).Return(newUserID, nil).Once()
			},
			want: domain.CreateResponse{
				ID: newUserID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepository: tt.fields.userRepository,
				hasher:         tt.fields.hasher,
			}
			tt.pre()
			got, err := uu.Create(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_DeleteByID(t *testing.T) {
	mockUserRepository := mockUserDomain.NewMockUserRepository(t)

	type fields struct {
		userRepository domain.UserRepository
	}
	type args struct {
		id uuid.UUID
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
				userRepository: mockUserRepository,
			},
			args: args{
				id: newUserID,
			},
			pre: func() {
				mockUserRepository.EXPECT().DeleteByID(newUserID).Return(nil).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepository: tt.fields.userRepository,
			}
			tt.pre()
			if err := uu.DeleteByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userUsecase_GetByID(t *testing.T) {
	mockUserRepository := mockUserDomain.NewMockUserRepository(t)

	type fields struct {
		userRepository domain.UserRepository
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		want    domain.GetByIDResponse
		wantErr bool
	}{
		{
			name: "get by id no result",
			fields: fields{
				userRepository: mockUserRepository,
			},
			args: args{
				id: newUserID,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByID(newUserID).Return(domain.User{}, sql.ErrNoRows).Once()
			},
			want:    domain.GetByIDResponse{},
			wantErr: true,
		},
		{
			name: "get by id error",
			fields: fields{
				userRepository: mockUserRepository,
			},
			args: args{
				id: newUserID,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByID(newUserID).Return(domain.User{}, errors.New("")).Once()
			},
			want:    domain.GetByIDResponse{},
			wantErr: true,
		},
		{
			name: "get by id error",
			fields: fields{
				userRepository: mockUserRepository,
			},
			args: args{
				id: newUserID,
			},
			pre: func() {
				mockUserRepository.EXPECT().GetByID(newUserID).Return(inputUser, nil).Once()
			},
			want: domain.GetByIDResponse{
				User: domain.User{
					Email: inputUser.Email,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepository: tt.fields.userRepository,
			}
			tt.pre()
			got, err := uu.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUsecase_UpdateRoleByID(t *testing.T) {
	mockUserRepository := mockUserDomain.NewMockUserRepository(t)

	type fields struct {
		userRepository domain.UserRepository
		hasher         hasher.Hasher
	}
	type args struct {
		id     uuid.UUID
		roleID constants.UserRole
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		pre     func()
		wantErr bool
	}{
		{
			name: "role id invalid",
			fields: fields{
				userRepository: mockUserRepository,
			},
			args: args{
				id:     newUserID,
				roleID: constants.UserRole(badRoleID),
			},
			pre: func() {
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				userRepository: mockUserRepository,
			},
			args: args{
				id:     newUserID,
				roleID: newRoleID,
			},
			pre: func() {
				mockUserRepository.EXPECT().UpdateRoleByID(newUserID, newRoleID).Return(nil).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uu := &userUsecase{
				userRepository: tt.fields.userRepository,
				hasher:         tt.fields.hasher,
			}
			tt.pre()
			if err := uu.UpdateRoleByID(tt.args.id, tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRoleByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
