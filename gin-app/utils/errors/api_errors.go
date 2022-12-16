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
