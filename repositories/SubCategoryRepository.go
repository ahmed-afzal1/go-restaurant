package repositories

import (
	"errors"

	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
)

func GetAllSubCategories() ([]models.SubCategory, error) {
	var subcategories []models.SubCategory

	result := config.DB.Find(&subcategories)

	if result.Error != nil {
		return []models.SubCategory{}, errors.New(result.Error.Error())
	}

	return subcategories, nil
}
