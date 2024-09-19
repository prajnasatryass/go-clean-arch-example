package apperror

import (
	"errors"
	"net/http"
)

type AppError struct {
	Status int
	Err    error
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func BadRequest(err error) error {
	return &AppError{
		Status: http.StatusBadRequest,
		Err:    err,
	}
}

func Unauthorized(err error) error {
	return &AppError{
		Status: http.StatusUnauthorized,
		Err:    err,
	}
}

func Forbidden(err error) error {
	return &AppError{
		Status: http.StatusForbidden,
		Err:    err,
	}
}

func NotFound(err error) error {
	return &AppError{
		Status: http.StatusNotFound,
		Err:    err,
	}
}

func InternalServerError(err error) error {
	return &AppError{
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}

func MethodNotImplemented() error {
	return &AppError{
		Status: http.StatusInternalServerError,
		Err:    errors.New("method not implemented"),
	}
}
