package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/prajnasatryass/go-clean-arch-example/internal/auth/domain"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/apperror"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/appresponse"
	"net/http"
)

type authController struct {
	authUsecase domain.AuthUsecase
}

func NewAuthController(eg *echo.Group, authUsecase domain.AuthUsecase) {
	c := &authController{
		authUsecase: authUsecase,
	}

	eg.POST("", c.Login)
	eg.PUT("/:refreshToken", c.Refresh)
	eg.DELETE("/:refreshToken", c.Logout)
}

func (ac *authController) Login(c echo.Context) error {
	var req domain.LoginRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	resp, err := ac.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.SuccessResponseBuilder(resp).Return(c, http.StatusCreated)
}

func (ac *authController) Refresh(c echo.Context) error {
	var req domain.RefreshRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	resp, err := ac.authUsecase.Refresh(req.RefreshToken)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.SuccessResponseBuilder(resp).Return(c)
}

func (ac *authController) Logout(c echo.Context) error {
	var req domain.LogoutRequest
	if err := c.Bind(&req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}
	if err := c.Validate(req); err != nil {
		return appresponse.ErrorResponseBuilder(apperror.BadRequest(err)).Return(c)
	}

	err := ac.authUsecase.Logout(req.RefreshToken)
	if err != nil {
		return appresponse.ErrorResponseBuilder(err).Return(c)
	}

	return appresponse.NoContent(c)
}
