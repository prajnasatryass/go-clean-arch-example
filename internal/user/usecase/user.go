package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/apperror"
	"github.com/prajnasatryass/tic-be/pkg/constants"
	"github.com/prajnasatryass/tic-be/pkg/hasher"
)

var (
	errEmailIsAlreadyUsed = func(email string) error {
		return errors.New(email + " is already used")
	}
	errHashPassword = func(err error) error { return errors.New("hash password error: " + err.Error()) }
	errCreateUser   = func(err error) error { return errors.New("create user error: " + err.Error()) }
	errUserNotFound = func(id uuid.UUID) error { return errors.New("user " + id.String() + " not found") }
	errGetUser      = func(err error) error { return errors.New("get user error: " + err.Error()) }
	errRoleNotFound = func(roleID constants.UserRole) error { return fmt.Errorf("role with ID %v not found", roleID) }
)

type userUsecase struct {
	userRepository domain.UserRepository
	hasher         hasher.Hasher
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		hasher:         hasher.NewHasher(),
	}
}

// TODO: UPDATE SWAGGER RESPONSE FROM THIS POINT
func (uu *userUsecase) Create(email, password string) (domain.CreateResponse, error) {
	existingUser, err := uu.userRepository.GetByEmail(email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return domain.CreateResponse{}, apperror.InternalServerError(err)
	}
	if existingUser.Email != "" {
		return domain.CreateResponse{}, apperror.Conflict(errEmailIsAlreadyUsed(email))
	}

	hashedPassword, err := uu.hasher.HashPassword(password)
	if err != nil {
		return domain.CreateResponse{}, apperror.InternalServerError(errHashPassword(err))
	}

	newUserID, err := uu.userRepository.Create(email, hashedPassword)
	if err != nil {
		return domain.CreateResponse{}, apperror.InternalServerError(errCreateUser(err))
	}

	return domain.CreateResponse{
		ID: newUserID,
	}, nil
}

func (uu *userUsecase) GetByID(id uuid.UUID) (domain.GetByIDResponse, error) {
	user, err := uu.userRepository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.GetByIDResponse{}, apperror.NotFound(errUserNotFound(id))
		}
		return domain.GetByIDResponse{}, apperror.InternalServerError(errGetUser(err))
	}
	user.Password = ""
	return domain.GetByIDResponse{
		User: user,
	}, nil
}

func (uu *userUsecase) UpdateRoleByID(id uuid.UUID, roleID constants.UserRole) error {
	if !roleID.Valid() {
		return apperror.NotFound(errRoleNotFound(roleID))
	}
	return uu.userRepository.UpdateRoleByID(id, roleID)
}

func (uu *userUsecase) DeleteByID(id uuid.UUID) error {
	return uu.userRepository.DeleteByID(id)
}
