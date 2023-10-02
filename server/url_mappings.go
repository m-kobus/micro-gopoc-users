package server

import (
	"micro-gopoc-users/handlers/health"
	"micro-gopoc-users/handlers/users"
)

func mapUrls() {
	router.GET("/check", health.HealthCheck)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/user", users.CreateUser)
	router.GET("/users/search", users.Search)
}
