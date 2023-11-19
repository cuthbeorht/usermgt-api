package main

import "usermgt-api/app/database"

func main() {
	db := database.Connect()

	database.Disconnect(db)
}
