package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect(host string, port int, user string, password string, dbname string) error {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		os.Exit(1)
	}
	fmt.Println("Connection: ", db)
	db.Ping()

	rows, err := db.Query("select * from users")

	if err != nil {
		fmt.Println("Error fetching users: ", err)
		os.Exit(1)
	}

	fmt.Println("Rows: ", rows)

	return nil
}
