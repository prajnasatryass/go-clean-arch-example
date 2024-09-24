package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/prajnasatryass/tic-be/internal/user/domain"
	"github.com/prajnasatryass/tic-be/pkg/apperror"
	"github.com/prajnasatryass/tic-be/pkg/appresponse"
	"net/http"
)

type userController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(eg *echo.Group, userUsecase domain.UserUsecase) {
	uc := &userController{
		userUsecase: userUsecase,
	}

	eg.POST("", uc.Create)
	eg.GET("/:id", uc.GetByID)
	eg.DELETE("/:id", uc.DeleteByID)
	return
}

func (uc *userController) Create(c echo.Context) error {
	var req domain.CreateRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	resp, err := uc.userUsecase.Create(req.Email, req.Password)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.SuccessResponseBuilder(resp).Return(c, http.StatusCreated)
}

func (uc *userController) GetByID(c echo.Context) error {
	var req domain.GetByIDRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	resp, err := uc.userUsecase.GetByID(req.ID)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.SuccessResponseBuilder(resp).Return(c)
}

func (uc *userController) DeleteByID(c echo.Context) error {
	var req domain.DeleteByIDRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	err := uc.userUsecase.DeleteByID(req.ID)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.NoContent(c)
}
