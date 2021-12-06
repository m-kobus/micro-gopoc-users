package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"date_created"`
}
