package user

import (
	"fmt"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	u := User{Id: 999, Email: "temp.user@server.com"}
	fmt.Println("user created: ", u)

	return c.JSON(200, &u)
}
