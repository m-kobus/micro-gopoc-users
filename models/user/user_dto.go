package users

import (
	"micro-gopoc-users/utils/errors"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateCreated string    `json:"date_created"`
}

func (u *User) Validate() *errors.RestError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("Empty string is not an email address")
	}
	return nil
}
