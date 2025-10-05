package controllers

import (
	models "TheAdidasTM/Models"
	service "TheAdidasTM/Service"
	"log"

	"github.com/gin-gonic/gin"
)

func RequestFromIlya(c *gin.Context) {
	var requestFromIlya models.RequestData
	if err := c.ShouldBindJSON(&requestFromIlya); err != nil {
		c.JSON(400, gin.H{
			"error": "bad request",
		})
		return
	}
	result, err := service.EventsProcess(requestFromIlya)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(result.Today)
	log.Println(result.Tommorow)
	c.JSON(200, gin.H{
		"today":    result.Today,
		"tomorrow": result.Tommorow,
	})
}
