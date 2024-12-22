package main

import (
	"github.com/Desk888/go-jwt/internal/controllers"
	"github.com/Desk888/go-jwt/internal/middleware"
	"github.com/Desk888/go-jwt/internal/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	// Initializers
	initializers.LoadEnv()
	initializers.InitDB()
	initializers.MigrateTables()
}

func main() {

	r := gin.Default()

	// Authentication Routes
	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}