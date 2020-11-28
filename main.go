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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Luke, you’re going to find that many of the truths we cling to depend greatly on our own point of view.",
		})
	})
	return
}
