package users

import (
	"fmt"
	"micro-gopoc-users/utils/dates"
	"micro-gopoc-users/utils/errors"

	"github.com/google/uuid"
)

// mocked database for dummies
var (
	usersDB = make(map[uuid.UUID]*User)
)

func (u *User) Get() *errors.RestError {
	r := usersDB[u.Id]
	if r == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User with id:%d not found", u.Id))
	}

	u.Id = r.Id
	u.FirstName = r.FirstName
	u.LastName = r.LastName
	u.Email = r.Email
	u.DateCreated = r.DateCreated

	return nil
}

func (u *User) Save() *errors.RestError {
	su := usersDB[u.Id]
	if su != nil {
		if su.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("User with email:%v already saved", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User with id:%d already saved", u.Id))
	}

	u.DateCreated = dates.GetCurrentDateTimeString()

	usersDB[u.Id] = u
	return nil
}
