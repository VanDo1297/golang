package errors

import (
	"net/http"

	"github.com/go-errors/errors"
)

type AppError struct {
	Err        *errors.Error
	StatusCode int
	Message    *string
}

func (appErr *AppError) Error() string {
	return appErr.Err.ErrorStack()
}

func WrapMessage(message string) *AppError {
	return WrapWithMessage(errors.New(message), message)
}

func Wrap(err error) *AppError {
	return WrapResponse(err, nil, http.StatusInternalServerError)
}

func WrapWithMessage(err error, message string) *AppError {
	return WrapResponse(err, &message, http.StatusBadRequest)
}

func WrapResponse(err error, message *string, statusCode int) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	if stackTraceError, ok := err.(*errors.Error); ok {
		return &AppError{
			Err:        stackTraceError,
			Message:    message,
			StatusCode: statusCode,
		}
	}
	return &AppError{
		Err:        errors.Wrap(err, 0),
		Message:    message,
		StatusCode: statusCode,
	}
}
