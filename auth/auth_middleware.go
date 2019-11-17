package auth

import (
	"bara/model"
	"bara/user"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/sessions"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}
var sessionCtxKey = &contextKey{"session"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(user user.RepositoryRunner, pool *redis.Pool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-token")

			cookie, err := url.QueryUnescape(c.Value)

			if err != nil {
				return
			}

			if cookie[0:2] != "s:" {
				return
			}

			redisKey := strings.Split(cookie[2:], ".")[0]

			session, err := getSession("sess:", redisKey, pool)

			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			userID, ok := session["passport"].(int64)
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			repo := user.GetRepository()
			// get the user from the database
			user, err := repo.GetUserByID(r.Context(), userID)

			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func getSession(prefix string, key string, pool *redis.Pool) (map[string]interface{}, error) {
	conn := pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, err
	}
	data, err := conn.Do("GET", prefix+key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil // no data was associated with this key
	}
	b, err := redis.Bytes(data, err)
	if err != nil {
		return nil, err
	}
	return deserialize(b), nil
}

func deserialize(d []byte) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal(d, &m)
	if err != nil {
		fmt.Printf("redistore.JSONSerializer.deserialize() Error: %v", err)
		return m
	}

	return m

}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.Users {
	raw, _ := ctx.Value(userCtxKey).(*model.Users)
	return raw
}

// Todo
func ForSessionContext(ctx context.Context) *sessions.CookieStore {
	raw, _ := ctx.Value(sessionCtxKey).(*sessions.CookieStore)
	return raw
}
