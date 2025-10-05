package controllers

import (
	models "TheAdidasTM/Models"
	service "TheAdidasTM/Service"

	"github.com/gin-gonic/gin"
)

var apiKey = "783e0858-39de-4c83-a72c-bc2858c795be"
var layout string = ""

func RequestFromIlya(c *gin.Context) {
	var requestFromIlya models.RequestData
	if err := c.ShouldBindJSON(&requestFromIlya); err != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}
	result, _ := service.EventsProcess(requestFromIlya)
	c.JSON(200, gin.H{
		"answer": result,
	})
}
