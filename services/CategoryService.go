package services

import (
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
	"github.com/ahmed-afzal1/restaurant/requests"
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
