package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type handler struct {
	DB *sql.DB
}

func New(db *sql.DB) handler {
	return handler{db}
}

type HealthResponse struct {
	Status string
}

func (h handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "OK"})
}
