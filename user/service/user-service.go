package service

import (
	"errors"
	"usermgt-api/user/models"
)

type CreateNewUser models.CreateUser

type UserService interface {
	create() (models.User, error)
}

func (nu CreateNewUser) create() (models.User, error) {

	if nu.Email == "" {
		return models.User{}, errors.New("new user is empty")
	}

	u := models.User{Id: 999, Email: nu.Email}

	return u, nil
}
