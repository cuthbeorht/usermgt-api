package repository

import (
	"usermgt-api/user/models"

	_ "github.com/lib/pq"
)

type Repository interface {
	persist() (models.User, error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "toor"
	db_name  = "users"
)

// func persist(u models.CreateUser) (models.User, error) {
// 	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)

// 	db, err := sql.Open("postgres", connString)

// 	if err != nil {
// 		panic(err)
// 	}

// 	db.Exec("insert into users (email) values ($1)", u.Email)
// }
