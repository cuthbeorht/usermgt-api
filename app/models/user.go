package models

type User struct {
	Id       string `json:"Id"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	FullName string `json:"FullName"`
}
