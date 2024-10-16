package services

import (
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
)

func GetAllSubCategories() ([]models.SubCategory, error) {
	subCategories, err := repositories.GetAllSubCategories()

	if err != nil {
		return []models.SubCategory{}, err
	}

	return subCategories, nil
}
