package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/persons", GeneratePersons)
	r.GET("/sentence", GenerateSentence)

	r.Run(":8083")
}
