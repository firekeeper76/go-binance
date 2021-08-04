package binance

import (
	"fmt"
	"net/http"
)

// APIError define API error with response header and http code
type APIError struct {
	http.Header
	HttpCode int
	Code     int64  `json:"code"`
	Message  string `json:"msg"`
}

func NewApiErr(message string) *APIError {
	return &APIError{
		Message: message,
	}
}

// Error return error code and message
func (e *APIError) Error() string {
	return fmt.Sprintf("<APIError> httpCode=%d, code=%d, msg=%s", e.HttpCode, e.Code, e.Message)
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
