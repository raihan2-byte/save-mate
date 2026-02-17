package util

import "github.com/go-playground/validator/v10"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func APIResponse(status int, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func APIError(status int, message string, errors interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Errors:  errors,
	}
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
