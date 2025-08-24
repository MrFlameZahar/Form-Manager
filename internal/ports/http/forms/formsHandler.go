package formsHandler

import (
	"FormManager/internal/model"
	authHandler "FormManager/internal/ports/http/auth"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

type FormInterface interface {
	CreateForm(form model.Form, user *model.User) (uint, error)
	EditForm(form model.Form) error
	DeleteForm(formID uint, user model.User) error
	GetForm(formID uint) (model.Form, error)
}

func CreateForm(form FormInterface, auth authHandler.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := auth.Me(tokenString)
		if err != nil {
			http.Error(w, "Failed to authenticate user: "+err.Error(), http.StatusUnauthorized)
			return
		}
		var userForm model.Form

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&userForm); err != nil {
			http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		formID, err := form.CreateForm(userForm, user)
		if err != nil {
			http.Error(w, "Create faild: "+err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, formID)
	}
}

func EditForm(form FormInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation goes here
	}
}

func DeleteForm(form FormInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation goes here
	}
}
func GetForm(form FormInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation goes here
	}
}
