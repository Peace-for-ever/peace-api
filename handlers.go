package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Peace-for-ever/peace-api/models"
)

type Param struct {
	MaxX       int `form:"maxX"`
	MaxY       int `form:"maxY"`
	MaxPersons int `form:"maxPersons"`
}

//GeneratePersonsHandler generate person
func GeneratePersonsHandler(c *gin.Context) {
	var param Param
	err := c.Bind(&param)
	checkError(c, err)

	var persons []models.Person

	if param.MaxPersons >= 50000 {
		c.JSON(406, gin.H{
			"message": "reste calme",
		})
		return
	}

	for i := 0; i < param.MaxPersons; i++ {
		persons = append(persons, models.RandomPerson(param.MaxX, param.MaxY))
	}
	c.JSON(http.StatusOK, persons)
}

//GenerateSentenceHandler generate sentence
func GenerateSentenceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, models.GenerateSentence())
}

func checkError(c *gin.Context, err error) {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(200, gin.H{"message": err.Error()})
	}
}
