package auth

import (
	"bara/model"
	"bara/user"
	"context"
	"net/http"

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
func Middleware(user user.RepositoryRunner, store *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "cookie-name")

			// if session == nil {
			// next.ServeHTTP(w, r)

			// r = r.WithContext(ctx)
			// next.ServeHTTP(w, r)
			// }

			if err != nil {
				http.Error(w, "Invalid cookie", http.StatusForbidden)
				return
			}

			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				ctx := context.WithValue(r.Context(), sessionCtxKey, session)
				r = r.WithContext(ctx)

				next.ServeHTTP(w, r)
				return
			}

			userID, ok := session.Values["userId"].(int64)
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
