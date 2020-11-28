package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run(":8083")
}

func setupRouter() (r *gin.Engine) {
	r = gin.Default()

	r.GET("/persons", GeneratePersonsHandler)
	r.GET("/sentence", GenerateSentenceHandler)
	return
}
