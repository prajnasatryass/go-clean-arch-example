package v1

import (
	"github.com/labstack/echo/v4"
	"tic-be/internal/user/domain"
)

type userController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(eg *echo.Group, userUsecase domain.UserUsecase) {
	return
}
