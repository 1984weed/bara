package endpoint

import (
	"bara"
	"bara/user"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type userEndpoint struct {
	uc    user.Usecase
	store *sessions.CookieStore
}

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

	session, _ := u.store.Get(r, bara.CookieAuthName)
	session.Values["authenticated"] = true
	session.Values["userID"] = me.ID

	session.Save(r, w)
}

type loginType struct {
	Email    string
	Password string
}

func (u *userEndpoint) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginInfo loginType
	err := decoder.Decode(&loginInfo)
	fmt.Println(err)
	if err != nil {
		return
	}

	me, err := u.uc.Login(r.Context(), "", loginInfo.Email, loginInfo.Password)

	if err != nil {
		return
	}

	session, _ := u.store.Get(r, bara.CookieAuthName)
	session.Values["authenticated"] = true
	session.Values["userID"] = me.ID

	session.Save(r, w)
}

func (u *userEndpoint) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := u.store.Get(r, bara.CookieAuthName)

	session.Values["authenticated"] = false
	session.Save(r, w)
}
