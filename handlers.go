package main

import (
	"fmt"
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
	Error(c, err)

	var persons []models.Person
	for i := 0; i < param.MaxPersons; i++ {
		persons = append(persons, models.RandomPerson(param.MaxX, param.MaxY))
	}

	fmt.Println(persons)
	c.JSON(http.StatusOK, persons)
}

func GenerateSentence(c *gin.Context) {
	c.JSON(http.StatusOK, models.GenerateSentence())
}

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(200, gin.H{"message": err.Error()})
		return true // signal that there was an error and the caller should return
	}
	return false // no error, can continue
}
