package usecase

import (
	"github.com/google/uuid"
	"github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/apperror"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uu *userUsecase) Create(email, password string) (domain.CreateResponse, error) {
	return domain.CreateResponse{}, apperror.MethodNotImplemented()
}

func (uu *userUsecase) GetByID(id uuid.UUID) (domain.GetByIDResponse, error) {
	return domain.GetByIDResponse{}, apperror.MethodNotImplemented()
}

func (uu *userUsecase) DeleteByID(id uuid.UUID) error {
	return apperror.MethodNotImplemented()
}
