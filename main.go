package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/persons", GeneratePersonsHandler)
	r.GET("/sentence", GenerateSentenceHandler)

	r.Run(":8083")
}
