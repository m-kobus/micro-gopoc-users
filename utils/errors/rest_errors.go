package errors

import "net/http"

type RestError struct {
	Message string `json:"message`
	Code    int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(m string) *RestError {
	return &RestError{
		Message: m,
		Code:    http.StatusBadRequest,
		Error:   "bad_request",
	}
}
