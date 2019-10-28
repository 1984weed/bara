package endpoint

import (
	"bara/user"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type userEndpoint struct {
	uc    user.Usecase
	store *sessions.CookieStore
}

const cookieAuthName = "auth"

// NewUserEndpoint creates user resolver
func NewUserEndpoint(uc user.Usecase, store *sessions.CookieStore) user.Endpoint {
	return &userEndpoint{
		uc,
		store,
	}
}

func (u *userEndpoint) SignUp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	email := r.FormValue("email")
	userName := r.FormValue("userName")
	password := r.FormValue("password")

	me, err := u.uc.Register(r.Context(), userName, email, password)
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	session, _ := u.store.Get(r, cookieAuthName)
	session.Values["authenticated"] = true
	session.Values["userID"] = me.ID

	session.Save(r, w)
}

func (u *userEndpoint) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	email := r.FormValue("email")
	userName := r.FormValue("userName")
	password := r.FormValue("password")
	session, _ := u.store.Get(r, cookieAuthName)

	me, err := u.uc.Login(r.Context(), userName, email, password)

	if err != nil {
		return
	}

	session.Values["authenticated"] = true
	session.Values["userID"] = me.ID

	session.Save(r, w)
}

func (u *userEndpoint) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := u.store.Get(r, cookieAuthName)

	session.Values["authenticated"] = false
	session.Save(r, w)
}
