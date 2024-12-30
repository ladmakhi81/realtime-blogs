package pkg_types

import (
	"fmt"
	"net/http"
)

type ClientError struct {
	StatusCode int
	Message    string
}

type ClientValidationError struct {
	ClientError
	Detail any
}

type ServerError struct {
	Message      string
	Method       string
	ErrorMessage string
}

func (richError *ServerError) Error() string {
	return fmt.Sprintf("error : %s", richError.Message)
}

func (richError *ClientError) Error() string {
	return fmt.Sprintf("error : %s", richError.Message)
}

func NewClientError(code int, mesaage string) *ClientError {
	return &ClientError{
		StatusCode: code,
		Message:    mesaage,
	}
}

func NewClientValidationError(detail any) *ClientValidationError {
	return &ClientValidationError{
		Detail: detail,
		ClientError: ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request body",
		},
	}
}

func NewServerError(mesaage, method, errorMessage string) *ServerError {
	return &ServerError{
		Message:      mesaage,
		Method:       method,
		ErrorMessage: errorMessage,
	}
}
