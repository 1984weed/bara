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

// JWTExpiredError represents that JWT is broken
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

func GraphqlPermissionError() error {
	return &gqlerror.Error{
		Message: "Current user doesn't have a permission",
	}

}
