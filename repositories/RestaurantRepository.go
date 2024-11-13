package repositories

import (
	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
)

func GetAllRestaurants() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	if err := config.DB.Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}
