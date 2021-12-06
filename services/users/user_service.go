package services

import (
	users "micro-gopoc-users/models/user"
	"micro-gopoc-users/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestError) {
	return &u, nil
}
