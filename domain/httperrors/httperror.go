package httperrors

import (
	"net/http"
)

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func NewInternalServeError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusInternalServerError,
		Error:   "Internal_server_Error",
	}
}

func NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}

/*func NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Message: message,
		Code:    http.StatusNotFound,
		Error:   "not_found",
	}
}*/
