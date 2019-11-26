package auth

import (
	"bara/model"
	"bara/user"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/garyburd/redigo/redis"
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
			c := r.Header.Get("Authorization")

			if c == "" {
				log.Print("Get auth-token cookie empty")
				next.ServeHTTP(w, r)
				return
			}
			cookie, err := url.QueryUnescape(c)

			if err != nil {
				log.Print("auth-token cookie is broken")
				next.ServeHTTP(w, r)
				return
			}

			if cookie[0:2] != "s:" {
				log.Print("auth-token cookie is broken")
				next.ServeHTTP(w, r)
				return
			}

			redisKey := strings.Split(cookie[2:], ".")[0]

			userID, err := getSession("sess:", redisKey, pool)

			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			repo := user.GetRepository()
			// get the user from the database
			user, err := repo.GetUserByID(r.Context(), userID)

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func getSession(prefix string, key string, pool *redis.Pool) (int64, error) {
	conn := pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return 0, err
	}
	data, err := conn.Do("GET", prefix+key)
	if err != nil {
		return 0, err
	}
	if data == nil {
		return 0, nil // no data was associated with this key
	}
	b, err := redis.Bytes(data, err)
	if err != nil {
		return 0, err
	}
	return getUserID(b), nil
}

func getUserID(d []byte) int64 {
	var f interface{}
	err := json.Unmarshal(d, &f)

	if err != nil {
		return 0
	}
	m := f.(map[string]interface{})

	user := m["passport"]

	v := user.(map[string]interface{})

	return int64(v["user"].(float64))
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.Users {
	raw, _ := ctx.Value(userCtxKey).(*model.Users)
	return raw
}
