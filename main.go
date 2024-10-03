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
	
	r.Run(":8080")
}


