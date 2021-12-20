package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error is the default error struct containing detailed information about the error
type Error struct {
	// HTTP Status code to be set in response
	Status int `json:"-"`
	// Message is the error message that may be displayed to end users
	Message string `json:"message,omitempty"`
}

var (
	// Generic error
	Generic = NewStatus(http.StatusInternalServerError)
	// DB represents database related errors
	DB = NewStatus(http.StatusInternalServerError)
	// Forbidden represents access to forbidden resource error
	Forbidden = NewStatus(http.StatusForbidden)
	// BadRequest represents error for bad requests
	BadRequest = NewStatus(http.StatusBadRequest)
	// NotFound represents errors for not found resources
	NotFound = NewStatus(http.StatusNotFound)
	// Unauthorized represents errors for unauthorized requests
	Unauthorized = NewStatus(http.StatusUnauthorized)
)

// NewStatus generates new error containing only http status code
func NewStatus(status int) *Error {
	return &Error{Status: status}
}

// New generates an application error
func New(status int, msg string) *Error {
	return &Error{Status: status, Message: msg}
}

// Error returns the error message.
func (e Error) Error() string {
	return e.Message
}

// Response writes an error response to client
func Response(c *gin.Context, err error) {
	switch err.(type) {
	case *Error:
		e := err.(*Error)
		if e.Message == "" {
			c.AbortWithStatus(e.Status)
		} else {
			c.AbortWithStatusJSON(e.Status, e)
		}
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
