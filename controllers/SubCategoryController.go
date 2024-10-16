package controllers

import (
	"net/http"

	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
)

// Details godoc
// @Summary      Get subcategories
// @Description  Returns the lists of subcategory
// @Tags         SubCategory
// @Security     JWT
// @Produce      json
// @Success      200 {object} map[string]interface{} "data: subcategories"
// @Failure      401 {object} map[string]interface{} "error: unauthorized"
// @Router       /admin/sub-category/index [get]
func GetAllSubCategories(c *gin.Context) {
	subCategories, err := services.GetAllSubCategories()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": subCategories})
}
