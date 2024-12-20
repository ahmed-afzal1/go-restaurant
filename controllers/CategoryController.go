package controllers

import (
	"context"
	"github.com/ahmed-afzal1/restaurant/helpers"
	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// Details godoc
// @Summary      Get categories
// @Description  Returns the lists of category
// @Tags         Category
// @Security     JWT
// @Produce      json
// @Success      200 {object} map[string]interface{} "data: categories"
// @Failure      401 {object} map[string]interface{} "error: unauthorized"
// @Router       /admin/category/index [get]
func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": categories})
}

// category store godoc
// @Summary      Category store
// @Description  Return a category after store
// @Tags         Category
// @Security     JWT
// @Produce      json
// @Param        login body requests.CategoryRequest true "Login Request"
// @Success      200 {object} map[string]interface{} "data: loginUser"
// @Failure      400 {object} map[string]interface{} "error: bad request"
// @Router       /admin/category/store [post]
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

// category store godoc
// @Summary      Category store
// @Description  Return a category after store
// @Tags         Category
// @Security     JWT
// @Produce      json
// @Param        id path string true "Category ID"
// @Success      200 {object} map[string]interface{} "data: loginUser"
// @Failure      400 {object} map[string]interface{} "error: bad request"
// @Router       /admin/category/edit/{id} [get]
func CategoryEdit(c *gin.Context) {
	id := c.Param("id")
	category, err := services.CategoryEdit(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": category})
}

func CategoryStatusUpdate(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")

	if id == "" || status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id or status not found"})
		return
	}

	category, err := services.CategoryStatusUpdate(id, status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": category})
}

func CategoryUpdate(c *gin.Context) {
	id := c.Param("id")
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

	category, err := services.CategoryUpdate(c, req, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": category})

}

func CategoryDelete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
		return
	}

	catDelMsg, err := services.CategoryDelete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": catDelMsg})
}
