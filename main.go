package main

import (
	"fmt"

	"usermgt-api/database"
	"usermgt-api/user/router"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Main")

	serve()
}

func serve() error {
	fmt.Println("Creating server")

	e := echo.New()

	database.Connect(
		"localhost",
		5432,
		"root",
		"toor",
		"users",
	)

	e.POST("/users", router.CreateUser)

	e.Logger.Fatal(e.Start(":9000"))

	return nil
}
