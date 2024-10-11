package repositories

import (
	"errors"

	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
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
