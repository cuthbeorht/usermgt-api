package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"usermgt-api/user/models"

	_ "github.com/lib/pq"
)

type Repository interface {
	Persist() (models.User, error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "toor"
	db_name  = "user"
)

type User models.CreateUser

func (u User) Persist() (models.User, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)

	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into users (email) values ($1)", u.Email)

	if err != nil {
		return models.User{}, nil
	}

	numRowsAffected, _ := result.RowsAffected()

	if numRowsAffected != 1 {
		return models.User{}, errors.New("error")
	}

	return models.User{Email: u.Email}, nil

}
