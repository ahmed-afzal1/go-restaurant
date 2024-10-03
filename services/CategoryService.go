package services

import (
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
)

func GetAllCategories() ([]models.Category, error) {
	categories, err := repositories.GetAllCategories()

	if err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}
