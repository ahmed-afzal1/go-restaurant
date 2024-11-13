package repositories

import (
	"errors"
	"os"

	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/helpers"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/gin-gonic/gin"
)

func GetCuisine() ([]models.Cuisine, error) {
	var cusines []models.Cuisine

	result := config.DB.Find(&cusines)

	if result.Error != nil {
		return []models.Cuisine{}, errors.New(result.Error.Error())
	}

	return cusines, nil
}

func CuisineStore(req requests.CuisineRequest, imagePath string) (models.Cuisine, error) {
	var cusine models.Cuisine

	cusine.Name = &req.Name
	cusine.Image = &imagePath

	result := config.DB.Create(&cusine)

	if result.Error != nil {
		return models.Cuisine{}, errors.New(result.Error.Error())
	}

	return cusine, nil
}

func CuisineEdit(id uint) (models.Cuisine, error) {
	var cusine models.Cuisine

	result := config.DB.First(&cusine, id)

	if result.Error != nil {
		return models.Cuisine{}, errors.New(result.Error.Error())
	}

	return cusine, nil
}

func CuisineUpdate(c *gin.Context, req requests.CuisineRequest, id string) (models.Cuisine, error) {
	var cuisine models.Cuisine

	result := config.DB.First(&cuisine, id)
	if result.Error != nil {
		return models.Cuisine{}, errors.New("category not found")
	}

	_, err := c.FormFile("image")
	if err == nil {
		if cuisine.Image != nil {
			err := os.Remove(*cuisine.Image)
			if err != nil && !os.IsNotExist(err) {
				return models.Cuisine{}, errors.New("failed to remove existing image")
			}
		}

		imagePath, err := helpers.ImageUploader(c)
		if err != nil {
			return models.Cuisine{}, err
		}
		cuisine.Image = &imagePath
	}

	cuisine.Name = &req.Name

	config.DB.Save(&cuisine)

	return cuisine, nil
}

func CuisineDelete(id string) (string, error) {
	var cuisine models.Cuisine

	result := config.DB.First(&cuisine, id)

	if result.Error != nil {
		return "", errors.New("cuisine not found")
	}

	config.DB.Delete(&cuisine)

	return "Cuisine deleted successfully", nil
}
