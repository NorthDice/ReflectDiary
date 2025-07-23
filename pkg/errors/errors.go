package apperr

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorType string

const (
	ErrorTypeValidation   ErrorType = "validation"
	ErrorTypeNotFound     ErrorType = "not_found"
	ErrorTypeUnauthorized ErrorType = "unauthorized"
	ErrorTypeForbidden    ErrorType = "forbidden"
	ErrorTypeConflict     ErrorType = "conflict"
	ErrorTypeInternal     ErrorType = "internal"
)

type AppError struct {
	Type        ErrorType `json:"type"`
	Message     string    `json:"message"`
	Code        string    `json:"code,omitempty"`
	Details     any       `json:"details,omitempty"`
	internalMsg string    `json:"-"`
	cause       error     `json:"-"`
}

func (e *AppError) Error() string {
	if e.internalMsg != "" {
		return e.internalMsg
	}
	return e.Message
}

func (e *AppError) GetHTTPStatus() int {
	switch e.Type {
	case ErrorTypeValidation:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrorTypeForbidden:
		return http.StatusForbidden
	case ErrorTypeConflict:
		return http.StatusConflict
	case ErrorTypeInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func (e *AppError) GetClientSafeError() *AppError {
	if e.Type == ErrorTypeInternal {

		return &AppError{
			Type:    ErrorTypeInternal,
			Message: "Something went wrong. Please try again later.",
			Code:    "INTERNAL_ERROR",
		}
	}

	return e
}

func (e *AppError) GetInternalDetails() string {
	details := fmt.Sprintf("Type: %s, Message: %s", e.Type, e.Message)
	if e.internalMsg != "" {
		details += fmt.Sprintf(", Internal: %s", e.internalMsg)
	}
	if e.cause != nil {
		details += fmt.Sprintf(", Cause: %v", e.cause)
	}
	return details
}

func NewValidationError(message string, details any) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Message: message,
		Code:    "VALIDATION_ERROR",
		Details: details,
	}
}

func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Type:    ErrorTypeNotFound,
		Message: fmt.Sprintf("%s not found", resource),
		Code:    "NOT_FOUND",
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeConflict,
		Message: message,
		Code:    "CONFLICT",
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeUnauthorized,
		Message: message,
		Code:    "UNAUTHORIZED",
	}
}

func NewInternalError(internalMsg string) *AppError {
	return &AppError{
		Type:        ErrorTypeInternal,
		Message:     "Something went wrong. Please try again later.",
		Code:        "INTERNAL_ERROR",
		internalMsg: internalMsg,
	}
}

func WrapInternalError(err error, context string) *AppError {
	return &AppError{
		Type:        ErrorTypeInternal,
		Message:     "Something went wrong. Please try again later.",
		Code:        "INTERNAL_ERROR",
		internalMsg: fmt.Sprintf("%s: %v", context, err),
		cause:       err,
	}
}

func IsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	ok := errors.As(err, &appErr)
	return appErr, ok
}
