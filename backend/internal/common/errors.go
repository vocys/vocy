package common

import (
	"net/http"
)

type AppError interface {
	error
	HttpStatusCode() int
}

type UiConfigDisabledError struct{}

func (e *UiConfigDisabledError) Error() string {
	return "The configuration can't be changed since the UI configuration is disabled"
}
func (e *UiConfigDisabledError) HttpStatusCode() int { return http.StatusForbidden }

type TooManyRequestsError struct{}

func (e *TooManyRequestsError) Error() string {
	return "Too many requests"
}
func (e *TooManyRequestsError) HttpStatusCode() int { return http.StatusTooManyRequests }
