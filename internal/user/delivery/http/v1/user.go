package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/prajnasatryass/go-clean-arch-example/internal/user/domain"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/apperror"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/appresponse"
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
	eg.PATCH("/:id/role", uc.UpdateRoleByID)
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
	var req domain.ByIDRequest
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

func (uc *userController) UpdateRoleByID(c echo.Context) error {
	var req domain.UpdateRoleByIDRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	err := uc.userUsecase.UpdateRoleByID(req.ID, req.RoleID)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.NoContent(c)
}

func (uc *userController) DeleteByID(c echo.Context) error {
	var req domain.ByIDRequest
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
