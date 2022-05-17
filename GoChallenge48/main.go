package main

import (
	"GoChallenge48/Go/public"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Tells to gin if we are in a dev environment or not
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// Tells to gin to force color in shell
	gin.ForceConsoleColor()

	router := gin.Default()

	router.Use(cors.Default())

	publicGroup := router.Group("/")
	{
		// Admin
		publicGroup.GET("/admin", public.GetAdmin)

		// Customers
		publicGroup.GET("/customer", public.GetCustomers)
		publicGroup.GET("/user/:id", public.GetUserById)

		//publicGroup.POST("/user", public.PostUser)

		// Organizer
		publicGroup.GET("/organizer", public.GetOrganizer)

		// Event
		publicGroup.GET("/event", public.GetEvents)
	}




	router.Run()
}

