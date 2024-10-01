package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"Status418/enums"
)
func main() {
	fmt.Println("Hello World!")
	var food enums.FoodType = 1
	fmt.Println(food)
	var moment enums.Moment= 0 
	fmt.Println(moment)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}


