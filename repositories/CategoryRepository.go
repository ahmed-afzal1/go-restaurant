package repositories

import (
	"errors"
	"os"
	"strconv"

	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/helpers"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/gin-gonic/gin"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	result := config.DB.Find(&categories)

	if result.Error != nil {
		return []models.Category{}, result.Error
	}

	return categories, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	saveData := config.DB.Create(&category)

	if saveData.Error != nil {
		return models.Category{}, errors.New("Failed to save category")
	}

	return category, nil
}

func CategoryEdit(id string) (models.Category, error) {
	var category models.Category
	result := config.DB.First(&category, id)

	if result.Error != nil {
		return models.Category{}, errors.New(result.Error.Error())
	}

	return category, nil
}

func CategoryStatusUpdate(id string, status string) (models.Category, error) {
	var category models.Category

	result := config.DB.First(&category)

	if result.Error != nil {
		return models.Category{}, errors.New(result.Error.Error())
	}

	statusUint, err := strconv.ParseUint(status, 10, 8)
	if err != nil {
		return models.Category{}, errors.New("invalid status value")
	}

	statusUint8 := uint8(statusUint)
	category.Status = &statusUint8
	config.DB.Save(&category)

	return category, nil
}

func CategoryUpdate(c *gin.Context, req requests.CategoryRequest, id string) (models.Category, error) {
	var category models.Category

	result := config.DB.First(&category, id)
	if result.Error != nil {
		return models.Category{}, errors.New("category not found")
	}

	_, err := c.FormFile("image")
	if err == nil {
		if category.Image != nil {
			err := os.Remove(*category.Image)
			if err != nil && !os.IsNotExist(err) {
				return models.Category{}, errors.New("failed to remove existing image")
			}
		}

		imagePath, err := helpers.ImageUploader(c)
		if err != nil {
			return models.Category{}, err
		}
		category.Image = &imagePath
	}

	category.Name = &req.Name
	category.Slug = &req.Name

	config.DB.Save(&category)

	return category, nil

}

func CategoryDelete(id string) (string, error) {
	var category models.Category

	result := config.DB.Delete(&category, id)

	if result.Error != nil {
		return "", errors.New(result.Error.Error())
	}

	return "Category remove successfully", nil
}
