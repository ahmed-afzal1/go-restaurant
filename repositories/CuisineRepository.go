package repositories

import (
	"errors"

	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
)

func GetCuisine() ([]models.Cuisine, error) {
	var cusines []models.Cuisine

	result := config.DB.Find(&cusines)

	if result.Error != nil {
		return []models.Cuisine{}, errors.New(result.Error.Error())
	}

	return cusines, nil
}
