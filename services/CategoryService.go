package services

import (
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/gin-gonic/gin"
)

func GetAllCategories() ([]models.Category, error) {
	categories, err := repositories.GetAllCategories()

	if err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}

func CategoryStore(req requests.CategoryRequest, imagePath string) (models.Category, error) {
	category := models.Category{
		Name:  &req.Name,
		Slug:  &req.Name,
		Image: &imagePath,
	}

	catRepo, err := repositories.CreateCategory(category)

	if err != nil {
		return models.Category{}, err
	}

	return catRepo, nil
}

func CategoryEdit(id string) (models.Category, error) {
	category, err := repositories.CategoryEdit(id)

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func CategoryStatusUpdate(id string, status string) (models.Category, error) {
	category, err := repositories.CategoryStatusUpdate(id, status)

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func CategoryUpdate(c *gin.Context, req requests.CategoryRequest, id string) (models.Category, error) {
	category, err := repositories.CategoryUpdate(c, req, id)

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func CategoryDelete(id string) (string, error) {
	delMsg, err := repositories.CategoryDelete(id)

	if err != nil {
		return "", err
	}

	return delMsg, nil
}
