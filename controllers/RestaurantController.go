package controllers

import (
	"net/http"

	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
)

func GetAllRestaurants(c *gin.Context) {
	resturants, err := services.GetAllRestaurants()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": resturants})
}

func RestaurantStore(c *gin.Context) {

}
