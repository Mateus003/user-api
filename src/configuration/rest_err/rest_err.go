package rest_err

import (
	"net/http"
)

type Rest_Err struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int    `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Rest_Err) Error() string{
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *Rest_Err {
	return &Rest_Err{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func BadRequestError(message string) *Rest_Err{
	return &Rest_Err{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Rest_Err{
	return &Rest_Err{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *Rest_Err{
	return &Rest_Err{
		Message: message,
		Err: "internal_error",
		Code: http.StatusInternalServerError,
	}
}