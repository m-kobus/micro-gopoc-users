package users

import (
	"net/http"

	users "micro-gopoc-users/models/user"
	services "micro-gopoc-users/services/users"
	errors "micro-gopoc-users/utils/errors"

	"github.com/gin-gonic/gin"
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

	// slower hand-made
	// b, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// if err := json.Unmarshal(b, &u) {
	// 	log.Println(err)
	// 	return
	// }

	r, err := services.CreateUser(u)
	if err != nil {
		c.JSON(err.Code, err)
	}
	c.JSON(http.StatusCreated, r)
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "to be continiue"})
}
