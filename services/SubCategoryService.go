package services

import (
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/repositories"
	"github.com/ahmed-afzal1/restaurant/requests"
)

func GetAllSubCategories() ([]models.SubCategory, error) {
	subCategories, err := repositories.GetAllSubCategories()

	if err != nil {
		return []models.SubCategory{}, err
	}

	return subCategories, nil
}

func SubCategoryStore(req requests.SubcategoryRequest) (models.SubCategory, error) {
	subcategory, err := repositories.SubCategoryStore(req)

	if err != nil {
		return models.SubCategory{}, err
	}

	return subcategory, nil
}

func SubCategoryEdit(id string) (models.SubCategory, error) {
	subcategory, err := repositories.SubCategoryEdit(id)

	if err != nil {
		return models.SubCategory{}, err
	}

	return subcategory, nil
}

func SubCategoryUpdate(req requests.SubcategoryRequest, id string) (models.SubCategory, error) {
	subcategory, err := repositories.SubCategoryUpdate(req, id)

	if err != nil {
		return models.SubCategory{}, err
	}

	return subcategory, nil
}

func SubCategoryDelete(id string) (string, error) {
	message, err := repositories.SubCategoryDelete(id)

	if err != nil {
		return "", err
	}

	return message, nil
}
