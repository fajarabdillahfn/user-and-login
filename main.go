package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	cDB "login-api/common/db/mongo"
	cMiddleware "login-api/common/middleware"

	cmdUser "login-api/cmd/user"
)

const port = ":8009"

func main() {
	log.Println("========== Start App ==========")
	initAllModules()
	router := mux.NewRouter()
	initRoutes(router)
	log.Println("========== Start Server ==========")
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Println(err)
	}
}

func initAllModules() {
	var ctx = context.Background()
	userDB, err := cDB.OpenDB(ctx)
	if err != nil {
		log.Println(err)
	}

	cmdUser.Initialize(userDB)
}

func initRoutes(r *mux.Router) {

	//this group doesn't need authentication
	noAuth := r.NewRoute().Subrouter()

	//LOGIN
	noAuth.HandleFunc("/api/v1/login", cmdUser.HTTPDelivery.Login).Methods("POST")

	//this group need bearer token authentication
	withAuth := r.NewRoute().Subrouter()
	withAuth.Use(cMiddleware.Middleware)

	//CREATE
	withAuth.HandleFunc("/api/v1/user", cmdUser.HTTPDelivery.CreateUser).Methods("POST")

	//RETRIEVE
	withAuth.HandleFunc("/api/v1/user", cmdUser.HTTPDelivery.GetUser).Methods("GET")

	//UPDATE
	withAuth.HandleFunc("/api/v1/user/{username}", cmdUser.HTTPDelivery.UpdateUser).Methods("PUT")

	//DELETE
	withAuth.HandleFunc("/api/v1/user/{username}", cmdUser.HTTPDelivery.DeleteUser).Methods("DELETE")
}
