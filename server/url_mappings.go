package server

import "micro-gopoc-users/controllers"

func mapUrls() {
	router.GET("/check", controllers.HealthCheck)
}
