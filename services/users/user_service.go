package services

import (
	users "micro-gopoc-users/models/user"
	"micro-gopoc-users/utils/errors"

	"github.com/google/uuid"
)

func CreateUser(u users.User) (*users.User, *errors.RestError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUser(uId uuid.UUID) (*users.User, *errors.RestError) {
	r := &users.User{Id: uId}
	if err := r.Get(); err != nil {
		return nil, err
	}
	return r, nil
}
