package controllers

import (
	"net/http"

	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": categories})
}
