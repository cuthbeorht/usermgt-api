package main

import (
	"log"
	"net/http"
	"usermgt-api/app/database"
	"usermgt-api/app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db := database.Connect()

	h := handlers.New(db)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(
		"/healthz",
		h.Health,
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/users",
		h.CreateUser,
	).Methods(http.MethodPost)

	router.HandleFunc("/users/{id}", h.UpdateUser).Methods(http.MethodPatch)

	log.Fatal(http.ListenAndServe(":9000", router))

	database.Disconnect(db)
}
