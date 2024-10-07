package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
	
)
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	
	r := gin.Default()
	foods := r.Group("/foods") 
	
	foods.GET("/:userId")
	foods.GET("/:userId/:code")
	foods.POST("/")
	foods.DELETE("/:code")
	foods.PUT("/:code")
	
	purchases := r.Group("/purchases")
	purchases.POST("/:userId")

	users:= r.Group("/users")
	users.GET("/")
	users.GET("/:id")
	users.POST("/")
	users.PUT("/:id")
	users.DELETE("/:id")	

	recipes := r.Group("/recipes")
	recipes.GET("/:userId")
	// recipes.GET("/:userId/:code")

	
	r.Run(":8080")
}


