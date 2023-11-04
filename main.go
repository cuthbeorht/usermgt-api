package main

import (
	"fmt"

	"usermgt-api/user"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Main")

	serve()
}

func serve() error {

	e := echo.New()

	e.POST("/users", user.CreateUser)

	e.Logger.Fatal(e.Start(":9000"))

	return nil
}