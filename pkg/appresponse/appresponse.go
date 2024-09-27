package appresponse

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/prajnasatryass/go-clean-arch-example/pkg/apperror"
	"net/http"
	"slices"
)

var httpSuccessStatusWhitelist = []int{
	http.StatusOK,
	http.StatusCreated,
	http.StatusAccepted,
	http.StatusNoContent,
	http.StatusResetContent,
}

var httpStatusDescriptionMap = map[int]string{
	http.StatusOK:                  "ok",
	http.StatusBadRequest:          "bad request",
	http.StatusUnauthorized:        "unauthorized",
	http.StatusForbidden:           "access forbidden",
	http.StatusNotFound:            "resource not found",
	http.StatusInternalServerError: "internal server error",
}

type baseResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	links   `json:",omitempty"`
}

type links struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}

type SuccessResponse struct {
	baseResponse `json:",inline"`
	Data         interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	baseResponse `json:",inline"`
	Error        string `json:"error"`
}

func SuccessResponseBuilder(data interface{}) SuccessResponse {
	return SuccessResponse{
		Data: data,
	}
}

func (resp SuccessResponse) Return(c echo.Context, httpStatus ...int) error {
	if len(httpStatus) > 0 && slices.Contains(httpSuccessStatusWhitelist, httpStatus[0]) {
		return c.JSON(httpStatus[0], resp)
	}
	return c.JSON(http.StatusOK, resp)
}

func ErrorResponseBuilder(err error) ErrorResponse {
	var appErr *apperror.AppError
	if errors.As(err, &appErr) {
		ae := err.(*apperror.AppError)

		return ErrorResponse{
			baseResponse: baseResponse{
				Status:  ae.Status,
				Message: httpStatusDescriptionMap[ae.Status],
			},
			Error: ae.Error(),
		}
	}

	errStr := "unknown error"
	if err != nil {
		errStr = err.Error()
	}

	return ErrorResponse{
		baseResponse: baseResponse{
			Status:  http.StatusInternalServerError,
			Message: httpStatusDescriptionMap[http.StatusInternalServerError],
		},
		Error: errStr,
	}
}

func (resp ErrorResponse) Return(c echo.Context) error {
	return c.JSON(resp.Status, resp)
}

func NoContent(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
