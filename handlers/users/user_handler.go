package users

import (
	"net/http"

	users "micro-gopoc-users/models/user"
	services "micro-gopoc-users/services/users"
	errors "micro-gopoc-users/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Search(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "to be continiue"})
}

func CreateUser(c *gin.Context) {
	u := users.User{}

	// fastest way
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid json body"))
		return
	}

	r, err := services.CreateUser(u)
	if err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusCreated, r)
}

func GetUser(c *gin.Context) {
	// func Parse(s string) (UUID, error)
	guid, perr := uuid.Parse(c.Param("user_id"))
	if perr != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Cannot parse given user_id"))
		return
	}

	r, err := services.GetUser(guid)
	if err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusOK, r)
}
