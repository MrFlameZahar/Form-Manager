package authHandler

import (
	"FormManager/internal/model"
	"net/http"

	"github.com/go-chi/render"
)

type Auth interface {
	SignupUser(username string, email string, password string) error
	LoginUser(email string, password string) (string, error)
	Me(token string) (*model.User, error)
}

func Login(auth Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		token, err := auth.LoginUser(r.FormValue("email"), r.FormValue("password"))
		if err != nil {
			http.Error(w, "Failed to login user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, token)
	}
}

func Signup(auth Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		err := auth.SignupUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
		if err != nil {
			http.Error(w, "Failed to signup user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func Me(auth Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		user, err := auth.Me(r.FormValue("token"))
		if err != nil {
			http.Error(w, "Failed to get user information: "+err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, user)
	}
}
