package errors

import "net/http"

type ApiErr struct {
	Message string `json: "message"`
	Status  int    `json: "status"`
	Error   string `json: "error"`
}

func NewBadRequestError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
