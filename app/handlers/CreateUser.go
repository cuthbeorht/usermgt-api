package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"usermgt-api/app/models"

	"github.com/google/uuid"
)

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(500)
		return
	}

	var user models.User
	json.Unmarshal(body, &user)

	user.Id = (uuid.New()).String()

	q := `INSERT INTO users (user_id, email, password, created_at, deleted_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, null) RETURNING user_id;`
	err = h.DB.QueryRow(q, &user.Id, &user.Email, &user.Password).Scan(&user.Id)

	if err != nil {
		log.Println("failed to execute query: ", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
