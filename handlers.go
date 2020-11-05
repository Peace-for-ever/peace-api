package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lacazethomas/peaceland/models"
)

type Param struct {
	MaxX       int `form:"maxX"`
	MaxY       int `form:"maxY"`
	MaxPersons int `form:"maxPersons"`
}

// GET /persons
func GeneratePersons(c *gin.Context) {
	var param Param
	err := c.Bind(&param)
	checkError(c, err)

	var persons []models.Person
	for i := 0; i < param.MaxPersons; i++ {
		persons = append(persons, models.RandomPerson(param.MaxX, param.MaxY))
	}

	c.JSON(http.StatusOK, persons)
}

func GenerateSentence(c *gin.Context) {
	c.JSON(http.StatusOK, models.GenerateSentence())
}

func checkError(c *gin.Context, err error) {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(200, gin.H{"message": err.Error()})
	}
}
