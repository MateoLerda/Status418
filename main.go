package main

import (
	"github.com/gin-gonic/gin"
	"log"
)
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	
	r.Run(":8080")
}


