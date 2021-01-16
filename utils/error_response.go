package utils

import (
	"net/http"

	"github.com/vektah/gqlparser/gqlerror"
)

// ResponseError represents the error of http response
type ResponseError struct {
	Status  int
	err     string
	Content responseErrorContent
}

type responseErrorContent struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

func (r *ResponseError) Error() string {
	return r.err
}

// JWTExpiredError represents that JWT is expired
func JWTExpiredError() *ResponseError {
	return &ResponseError{
		Status: http.StatusUnauthorized,
		err:    "JWT token is expired",
		Content: responseErrorContent{
			Kind:    "JWT_EXPIRED",
			Message: "JWT token is expired",
		},
	}
}

// JWTBrokenError represents that JWT is broken
func JWTBrokenError(err string) *ResponseError {
	return &ResponseError{
		Status: http.StatusUnauthorized,
		err:    err,
		Content: responseErrorContent{
			Kind:    "JWT_EXPIRED",
			Message: "JWT token is broken",
		},
	}
}

func errorCode(code string) map[string]interface{} {
	return map[string]interface{}{
		"Code": code,
	}
}

// InternalServerError represents some unspecific errors
var (
	InternalServerError = &gqlerror.Error{
		Message:    "An internal server error occurred",
		Extensions: errorCode("INTERNAL_SERVER_ERROR"),
	}
	// PermissionError represents some unspecific errors
	PermissionError = &gqlerror.Error{
		Message:    "Current user doesn't have a permission",
		Extensions: errorCode("PERMISSION_ERROR"),
	}
)
