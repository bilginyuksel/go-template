package errors

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// HTTPErrorMapperFunc is a function that maps native go error to HTTPError
type HTTPErrorMapperFunc func(err error) *HTTPError

// HTTPError is a custom error type that is used to send error response to client
type HTTPError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

// NewHTTPError creates a new HTTPError
func NewHTTPError(statusCode int, message string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// Error implements error interface
func (he *HTTPError) Error() string {
	return he.Message
}

// EchoErrorHandler is a custom error handler written for echo http framework
// It converts native go error to HTTPError and writes it to response
func EchoErrorHandler(defaultErr *HTTPError, mappers ...HTTPErrorMapperFunc) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if herror, ok := err.(*HTTPError); ok {
			_ = c.JSON(herror.StatusCode, herror)
			return
		}

		for _, mapper := range mappers {
			if herror := mapper(err); herror != nil {
				zap.L().Error("http error mapper func mapped the error", zap.Error(err))
				_ = c.JSON(herror.StatusCode, herror)
				return
			}
		}

		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}
