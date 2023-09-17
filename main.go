package main

import (
	"api/controllers"
	"api/initializers"
	"api/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("Hello")
	fmt.Println("Hello2")
	fmt.Println(
		fmt.Printf("Application using database user: %s", os.Getenv("DB_USER")),
	)
	// start the actual application
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
