package services

import (
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
)

func GetAllRestaurants() ([]models.Restaurant, error) {
	repositories, err := repositories.GetAllRestaurants()

	if err != nil {
		return []models.Restaurant{}, err
	}

	return repositories, nil
}
