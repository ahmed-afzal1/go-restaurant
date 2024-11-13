package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := requests.FormatValidationError(validationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": formattedErrors})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	cusine, err := services.CuisineStore(c, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": cusine})

}

func CuisineEdit(c *gin.Context) {
	id := c.Param("id")
	cusine, err := services.CuisineEdit(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": cusine})
}

func CuisineUpdate(c *gin.Context) {

	var req requests.CuisineRequest
	id := c.Param("id")

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.ShouldBind(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := requests.FormatValidationError(validationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": formattedErrors})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	cusine, err := services.CuisineUpdate(c, req, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": cusine})

}

func CuisineDelete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
		return
	}

	cusineDelMsg, err := services.CuisineDelete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": cusineDelMsg})
}
