package main

import (
	_ "Back_End/docs"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	register(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
