package rest_err

import (
	"net/http"
)

type Resterr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Resterr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *Resterr {
	return &Resterr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *Resterr {
	return &Resterr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Resterr {
	return &Resterr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *Resterr {
	return &Resterr{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Resterr {
	return &Resterr{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Resterr {
	return &Resterr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
