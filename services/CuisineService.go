package services

import (
	"github.com/ahmed-afzal1/restaurant/helpers"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCuisine() ([]models.Cuisine, error) {
	cuisine, err := repositories.GetCuisine()

	if err != nil {
		return []models.Cuisine{}, err
	}

	return cuisine, nil
}

func CuisineStore(c *gin.Context, req requests.CuisineRequest) (models.Cuisine, error) {
	imagePath, err := helpers.ImageUploader(c)

	if err != nil {
		return models.Cuisine{}, err
	}

	cusine, err := repositories.CuisineStore(req, imagePath)

	if err != nil {
		return models.Cuisine{}, err
	}

	return cusine, nil
}

func CuisineEdit(id string) (models.Cuisine, error) {
	cusine_id, _ := strconv.Atoi(id)
	cusineId := uint(cusine_id)

	cusine, err := repositories.CuisineEdit(cusineId)

	if err != nil {
		return models.Cuisine{}, err
	}

	return cusine, nil
}

func CuisineUpdate(c *gin.Context, req requests.CuisineRequest, id string) (models.Cuisine, error) {
	cusine, err := repositories.CuisineUpdate(c, req, id)

	if err != nil {
		return models.Cuisine{}, err
	}

	return cusine, nil
}

func CuisineDelete(id string) (string, error) {
	cusineDelMsg, err := repositories.CuisineDelete(id)

	if err != nil {
		return "", err
	}

	return cusineDelMsg, nil
}
