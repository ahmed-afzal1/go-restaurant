package repositories

import (
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
