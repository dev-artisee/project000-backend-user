package cerror

import (
	"net/http"

	"github.com/pkg/errors"

	"project000-backend-user/pkg/response"
)

type (
	CustomError struct {
		Err           error
		Status        int
		Code          response.Code
		Message       string
		EnableLogging bool
		EnableAlert   bool
	}
)

func (e CustomError) Error() string {
	return e.Err.Error()
}

func (e CustomError) AppendMessage(message string) *CustomError {
	if len(message) > 0 {
		e.Message = e.Message + ": " + message
	}

	return &e
}

func NewCustomError(err error, status int, code response.Code, message string, enableLogging bool, enableAlert bool) *CustomError {
	var cause error
	if err != nil {
		cause = errors.Wrap(err, message)
	} else {
		cause = errors.New(message)
	}

	return &CustomError{
		Err:           cause,
		Status:        status,
		Code:          code,
		Message:       message,
		EnableLogging: enableLogging,
		EnableAlert:   enableAlert,
	}
}

func NewErrBadRequest(err error) *CustomError {
	return NewCustomError(
		err,
		http.StatusBadRequest,
		response.CodeBadRequest,
		"Bad Request",
		false,
		false,
	)
}

func NewErrUnauthorized(err error) *CustomError {
	return NewCustomError(
		err,
		http.StatusUnauthorized,
		response.CodeUnauthorized,
		"Unauthorized",
		false,
		false,
	)
}

func NewErrForbidden(err error) *CustomError {
	return NewCustomError(
		err,
		http.StatusForbidden,
		response.CodeForbidden,
		"Forbidden",
		false,
		false,
	)
}

func NewErrNotFound(err error) *CustomError {
	return NewCustomError(
		err,
		http.StatusNotFound,
		response.CodeNotFound,
		"Not Found",
		false,
		false,
	)
}

func NewErrInternalServerError(err error) *CustomError {
	return NewCustomError(
		err,
		http.StatusInternalServerError,
		response.CodeInternalServerError,
		"Internal Server Error",
		true,
		false,
	)
}
