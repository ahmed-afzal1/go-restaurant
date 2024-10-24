package services

import (
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
)

func GetCuisine() ([]models.Cuisine, error) {
	cuisine, err := repositories.GetCuisine()

	if err != nil {
		return []models.Cuisine{}, err
	}

	return cuisine, nil
}
