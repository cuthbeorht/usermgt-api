package models

type User struct {
	Id    int    `json:"userId"`
	Email string `json:"userEmail"`
}

type CreateUser struct {
	Email string `json:"userEmail"`
}
