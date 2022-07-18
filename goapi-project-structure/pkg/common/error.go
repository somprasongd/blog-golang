package common

import (
	"net/http"
	"strings"
)

type AppError struct {
	Code    int           `json:"code" example:"400"`
	Message string        `json:"message" example:"Invalid json body"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorDetail struct {
	Target  string `json:"target" example:"Name"`
	Message string `json:"message" example:"Name field is required"`
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) Is(target error) bool {
	return target.Error() == e.Message
}

func NewBadRequestError(message string) error {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInvalidError(details string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: "invalid data see details",
		Details: parseError(details),
	}
}

func NewUnauthorizedError(message string) error {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewForbiddenError(message string) error {
	return AppError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func parseError(err string) []ErrorDetail {
	eds := []ErrorDetail{}
	errs := strings.Split(err, ",")
	for _, e := range errs {
		kv := strings.Split(e, ":")
		eds = append(eds, ErrorDetail{
			Target:  strings.TrimSpace(kv[0]),
			Message: strings.TrimSpace(kv[1]),
		})
	}
	return eds
}
