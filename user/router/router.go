package router

import (
	"fmt"

	"usermgt-api/user/models"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	u := models.User{Id: 999, Email: "temp.user@server.com"}
	fmt.Println("user created: ", u)

	return c.JSON(200, &u)
}
