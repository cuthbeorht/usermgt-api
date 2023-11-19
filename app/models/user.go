package models

import "time"

type User struct {
	Id        string    `json:"Id"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	FullName  string    `json:"FullName"`
	CreatedAt time.Time `json:"CreatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
}
