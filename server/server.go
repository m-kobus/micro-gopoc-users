package server

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartServer() {
	mapUrls()
	router.Run(":8080")
}
