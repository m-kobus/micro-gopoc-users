package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"micro-gopoc-users/services/users"
)

func Search(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "to be continiue"})
}

func CreateUser(c *gin.Context) {
	u := user.User{}
	
	// fastest way
	if err := c.ShouldBindJSON(&u)
	
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
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, r)
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "to be continiue"})
}
