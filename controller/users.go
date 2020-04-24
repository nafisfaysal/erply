package controller

import (
	"github.com/nafisfaysal/erply/errors"
	"github.com/nafisfaysal/erply/models"
	"net/http"
)

type Users struct {
	models.UserService
}

func (u *Users) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	user := models.User{
		Email:    body.Email,
		Password: body.Password,
	}
	if err := u.UserService.Create(&user); err != nil {
		errors.ServeInternalServerError(w, err)
		return
	}
	err := saveToSession(r, w, "userID", user.ID)
	if err != nil {
		errors.ServeInternalServerError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (u *Users) HandleLogin(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	user, err := u.UserService.Authenticate(body.Email, body.Password)
	if err != nil {
		errors.AuthenticationFailed(w, err)
		return
	}

	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = saveToSession(r, w, "userID", user.ID)
	if err != nil {
		errors.ServeInternalServerError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)

}

func (u *Users) HandleLogout(w http.ResponseWriter, r *http.Request) {
	seesion, _ := store.Get(r, "erply")
	delete(seesion.Values, "userID")
	seesion.Options.MaxAge = -1
	_ = seesion.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
