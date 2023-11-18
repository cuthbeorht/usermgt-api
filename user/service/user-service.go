package service

import (
	"errors"
	"usermgt-api/user/models"
	"usermgt-api/user/repository"
)

type CreateNewUser models.CreateUser

type UserService interface {
	create() (models.User, error)
}

func (nu CreateNewUser) create() (models.User, error) {

	if nu.Email == "" {
		return models.User{}, errors.New("new user is empty")
	}

	newUser := models.CreateUser{Email: nu.Email}

	repository.Repository.Persist(newUser)

	if err != nil {
		return models.User{}, errors.New("unable to persist user")
	}

	return u, nil
}
