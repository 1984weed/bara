package auth

import (
	"bara/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var UserCtxKey = &contextKey{"user"}
var SessionCtxKey = &contextKey{"session"}

type contextKey struct {
	name string
}

// CurrentUser contains current's userId
type CurrentUser struct {
	Sub int64
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := r.Header.Get("Authorization")

			splittedToken := strings.Split(c, " ")

			// Guest user doesn't have Authorization token or empty token
			if len(splittedToken) != 2 || splittedToken[1] == "" {
				next.ServeHTTP(w, r)
				return

			}

			tokenString := splittedToken[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(jwtSecret), nil
			})
			// JWT token is broken
			if err != nil {
				interceptRequest(w, utils.JWTBrokenError(err.Error()))
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// put it in context
				ctx := context.WithValue(r.Context(), UserCtxKey, &CurrentUser{Sub: int64(claims["sub"].(float64))})

				// and call the next with our new context
				r = r.WithContext(ctx)
			} else {
				interceptRequest(w, utils.JWTExpiredError())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func interceptRequest(w http.ResponseWriter, res *utils.ResponseError) {
	w.WriteHeader(res.Status)
	json.NewEncoder(w).Encode(res.Content)
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *CurrentUser {
	raw, _ := ctx.Value(UserCtxKey).(*CurrentUser)
	return raw
}
