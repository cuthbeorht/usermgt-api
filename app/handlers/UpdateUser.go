package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"usermgt-api/app/models"

	"github.com/gorilla/mux"
)

type Error struct {
	Message string
}

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	queryParams := mux.Vars(r)
	userId := queryParams["id"]

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	var user models.User
	var foundUser models.User
	json.Unmarshal(body, &user)

	q := `SELECT * FROM users WHERE user_id = $1`
	err = h.DB.QueryRow(q, &userId).Scan(&foundUser.Id, &foundUser.Email, &foundUser.Password, &foundUser.CreatedAt, &foundUser.DeletedAt, &foundUser.FullName)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Error{Message: "User not found"})
	}

	fmt.Printf("Found user: %s", foundUser)

}
