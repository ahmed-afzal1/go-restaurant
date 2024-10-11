package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ahmed-afzal1/restaurant/helpers"
	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": categories})
}

func CategoryStore(c *gin.Context) {
	var req requests.CategoryRequest

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

	imagePath, err := helpers.ImageUploader(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	saveCategory, err := services.CategoryStore(req, imagePath)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": saveCategory})
}

func CategoryEdit(c *gin.Context) {
	id := c.Param("id")
	category, err := services.CategoryEdit(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": category})
}

func CategoryUpdate(c *gin.Context) {

}

func CategoryDelete(c *gin.Context) {

}
