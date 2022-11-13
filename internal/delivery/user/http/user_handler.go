package http

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	cWrapper "login-api/common/wrapper"
	"login-api/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

// Create Handler
func (d *Delivery) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Header.Get("accessRole") != "admin" {
		cWrapper.ErrorJSON(w, errors.New("unauthorized, admin only"), http.StatusUnauthorized)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	role := r.FormValue("role")

	if role == "" || username == "" || password == "" {
		cWrapper.ErrorJSON(w, errors.New("missing parameter{s}"), http.StatusBadRequest)
		return
	}

	if role != "admin" && role != "user" {
		cWrapper.ErrorJSON(w, errors.New("invalid role parameter"), http.StatusBadRequest)
		return
	}

	userCheck, _ := d.userUC.UserGetByUsername(ctx, username)
	if userCheck.Username != "" {
		cWrapper.ErrorJSON(w, errors.New("username already used"), http.StatusBadRequest)
		return
	}

	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	err := d.userUC.UserCreate(ctx, models.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	})
	if err != nil {
		log.Println(err)
		cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, fmt.Sprintf("user %s created", username), "message")
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}
}

// Read Handler
func (d *Delivery) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Header.Get("accessRole") == "user" {
		user, err := d.userUC.UserGetByUsername(ctx, r.Header.Get("accessUsername"))
		if err != nil {
			cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		err = cWrapper.WriteJSON(w, http.StatusOK, user, "data")
		if err != nil {
			cWrapper.ErrorJSON(w, err)
			return
		}
	} else {
		users, err := d.userUC.GetUserAll(ctx)
		if err != nil {
			cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}

		err = cWrapper.WriteJSON(w, http.StatusOK, users, "data")
		if err != nil {
			cWrapper.ErrorJSON(w, err)
			return
		}
	}
}

// Update Handler
func (d *Delivery) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Header.Get("accessRole") != "admin" {
		cWrapper.ErrorJSON(w, errors.New("unauthorized, admin only"), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	password := r.FormValue("password")
	role := r.FormValue("role")

	if username == "" {
		cWrapper.ErrorJSON(w, errors.New("missing username parameter"), http.StatusBadRequest)
		return
	}

	user, err := d.userUC.UserGetByUsername(ctx, username)
	if err != nil {
		if user.Username == "" {
			cWrapper.ErrorJSON(w, errors.New("user not found"), http.StatusNotFound)
			return
		}
		cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if role == "" {
		role = user.Role
	}

	if role != "admin" && role != "user" {
		cWrapper.ErrorJSON(w, errors.New("invalid role parameter"), http.StatusBadRequest)
		return
	}

	hashedPassword := ""
	if password == "" {
		hashedPassword = user.Password
	} else {
		hashedPassword = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	}

	updatedUser := models.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}

	err = d.userUC.UserUpdate(ctx, updatedUser)
	if err != nil {
		log.Println(err)
		cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, fmt.Sprintf("user %s updated", username), "message")
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}
}

// Delete Handler
func (d *Delivery) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Header.Get("accessRole") != "admin" {
		cWrapper.ErrorJSON(w, errors.New("unauthorized, admin only"), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	if username == "" {
		cWrapper.ErrorJSON(w, errors.New("missing username parameter"), http.StatusBadRequest)
		return
	}

	_, err := d.userUC.UserGetByUsername(ctx, username)
	if err != nil {
		cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = d.userUC.UserDelete(ctx, username)
	if err != nil {
		log.Println(err)
		cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, fmt.Sprintf("user %s deleted", username), "message")
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}
}
