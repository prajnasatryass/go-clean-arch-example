package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tic-be/internal/auth/domain"
	"tic-be/pkg/apperror"
	"tic-be/pkg/appresponse"
)

type authController struct {
	authUsecase domain.AuthUsecase
}

func NewAuthController(eg *echo.Group, authUsecase domain.AuthUsecase) {
	c := &authController{
		authUsecase: authUsecase,
	}

	eg.POST("/auth", c.Login)
	eg.PUT("/auth", c.Refresh)
	eg.DELETE("/auth", c.Logout)
}

func (uc *authController) Login(c echo.Context) error {
	var req domain.LoginRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	resp, err := uc.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.SuccessResponseBuilder(resp).Return(c, http.StatusCreated)
}

func (uc *authController) Refresh(c echo.Context) error {
	var req domain.RefreshRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	resp, err := uc.authUsecase.Refresh(req.RefreshToken)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.SuccessResponseBuilder(resp).Return(c)
}

func (uc *authController) Logout(c echo.Context) error {
	var req domain.LogoutRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	err := uc.authUsecase.Logout(req.RefreshToken)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.NoContent(c)
}
