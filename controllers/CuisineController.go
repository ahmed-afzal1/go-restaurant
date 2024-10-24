package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
)

// cuisine godoc
// @Summary      Get cuisinea
// @Description  Returns the lists of cuisine
// @Tags         Cuisine
// @Security     JWT
// @Produce      json
// @Success      200 {object} map[string]interface{} "data: cuisines"
// @Failure      401 {object} map[string]interface{} "error: unauthorized"
// @Router       /admin/cuisine/index [get]
func GetCuisine(c *gin.Context) {
	cusines, err := services.GetCuisine()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": cusines})
}

func CuisineStore(c *gin.Context) {
	var req requests.CuisineRequest

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.ShouldBind(&req); err != nil {

	}

}
