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

// Details godoc
// @Summary      Get subcategories
// @Description  Returns the lists of subcategory
// @Tags         SubCategory
// @Security     JWT
// @Produce      json
// @Param        subcategory body requests.SubcategoryRequest true "SubCategory Request"
// @Success      200 {object} map[string]interface{} "data: subcategories"
// @Failure      401 {object} map[string]interface{} "error: unauthorized"
// @Router       /admin/sub-category/store [post]
func SubCategoryStore(c *gin.Context) {
	var req requests.SubcategoryRequest

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

	subCategory, err := services.SubCategoryStore(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": subCategory})
}

// subCategory  godoc
// @Summary      Get SubCategory
// @Description  Return a category after store
// @Tags         SubCategory
// @Security     JWT
// @Produce      json
// @Param        id path string true "SubCategory ID"
// @Success      200 {object} map[string]interface{} "data: SubCategory"
// @Failure      400 {object} map[string]interface{} "error: bad request"
// @Router       /admin/sub-category/edit/{id} [get]
func SubCategoryEdit(c *gin.Context) {
	id := c.Param("id")
	subcategory, err := services.SubCategoryEdit(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": subcategory})
}

// Update subcategory godoc
// @Summary      Update subcategory
// @Description  Update the subcategory
// @Tags         SubCategory
// @Security     JWT
// @Produce      json
// @Param        subcategory body requests.SubcategoryRequest true "SubCategory Request"
// @Success      200 {object} map[string]interface{} "data: subcategories"
// @Failure      401 {object} map[string]interface{} "error: unauthorized"
// @Router       /admin/sub-category/update/{id} [patch]
func SubCategoryUpdate(c *gin.Context) {
	var req requests.SubcategoryRequest

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

	subCategory, err := services.SubCategoryStore(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": subCategory})
}

// subCategory delete  godoc
// @Summary      Delete SubCategory
// @Description  Delete SubCategory
// @Tags         SubCategory
// @Security     JWT
// @Produce      json
// @Param        id path string true "SubCategory ID"
// @Success      200 {object} map[string]interface{} "data: SubCategory"
// @Failure      400 {object} map[string]interface{} "error: bad request"
// @Router       /admin/sub-category/delete/{id} [delete]
func SubCategoryDelete(c *gin.Context) {
	id := c.Param("id")

	message, err := services.CategoryDelete(id)

	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "data": message})

}
