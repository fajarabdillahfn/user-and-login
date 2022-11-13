package http

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"time"

	cWrapper "login-api/common/wrapper"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("deall")

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type Response struct {
	Token string `json:"access_token"`
}

func (d *Delivery) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		cWrapper.ErrorJSON(w, errors.New("missing username or password"), http.StatusBadRequest)
		return
	}

	user, err := d.userUC.UserGetByUsername(ctx, username)
	if err != nil {
		cWrapper.ErrorJSON(w, errors.New("invalid username or password"), http.StatusUnauthorized)
		return
	}

	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	if hashedPassword != user.Password {
		cWrapper.ErrorJSON(w, errors.New("invalid username or password"), http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(25 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		cWrapper.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, Response{Token: tokenString}, "token")
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}
}
