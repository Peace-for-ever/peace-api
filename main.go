package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/persons", GeneratePersonsHandler)
	r.GET("/sentence", GenerateSentenceHandler)
	return
}

func main() {
	r := setupRouter()
	r.Run(":8083")
}
