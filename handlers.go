package main

import (
	"net/http"

	"github.com/Peace-for-ever/peace-api/models"
	"github.com/gin-gonic/gin"
)

//Param handlers
type Param struct {
	Amount int `form:"amount"`
}

//GeneratePersonsHandler generate person
func GeneratePersonsHandler(c *gin.Context) {
	var param Param
	err := c.Bind(&param)
	checkError(c, err)

	var events []models.Drone

	if param.Amount >= 500000 {
		c.JSON(406, gin.H{
			"message": "reste calme",
		})
		return
	}

	for i := 0; i < param.Amount; i++ {
		events = append(events, models.GenerateEvent())
	}
	c.JSON(http.StatusOK, events)
}

func checkError(c *gin.Context, err error) {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(200, gin.H{"message": err.Error()})
	}
}
