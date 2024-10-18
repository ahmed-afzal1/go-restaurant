package repositories

import (
	"errors"
	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/requests"
	"strconv"
)

func GetAllSubCategories() ([]models.SubCategory, error) {
	var subcategories []models.SubCategory

	result := config.DB.Find(&subcategories)

	if result.Error != nil {
		return []models.SubCategory{}, errors.New(result.Error.Error())
	}

	return subcategories, nil
}

func SubCategoryStore(req requests.SubcategoryRequest) (models.SubCategory, error) {
	var category models.Category
	var subCategory models.SubCategory

	categoryResult := config.DB.First(&category, req.CategoryId)

	if categoryResult.Error != nil {
		return models.SubCategory{}, errors.New("Category not found!")
	}

	cat_id, _ := strconv.Atoi(req.CategoryId)
	catId := uint(cat_id)

	subCategory.Name = &req.Name
	subCategory.CategoryID = catId

	result := config.DB.Create(&subCategory)

	if result.Error != nil {
		return models.SubCategory{}, errors.New(result.Error.Error())
	}

	return subCategory, nil
}

func SubCategoryEdit(id string) (models.SubCategory, error) {
	var subcategory models.SubCategory

	result := config.DB.Find(&subcategory, id)

	if result.Error != nil {
		return models.SubCategory{}, errors.New("Subcategory not found")
	}

	return subcategory, nil
}

func SubCategoryUpdate(req requests.SubcategoryRequest, id string) (models.SubCategory, error) {
	var category models.Category
	var subcategory models.SubCategory

	subcategoryResult := config.DB.Find(&subcategory, req.CategoryId)

	if subcategoryResult.Error != nil {
		return models.SubCategory{}, errors.New("subcategory not found")
	}

	categoryResult := config.DB.First(&category, req.CategoryId)

	if categoryResult.Error != nil {
		return models.SubCategory{}, errors.New("Category not found!")
	}

	cat_id, _ := strconv.Atoi(req.CategoryId)
	catId := uint(cat_id)

	subcategory.Name = &req.Name
	subcategory.CategoryID = catId

	result := config.DB.Save(&subcategory)

	if result.Error != nil {
		return models.SubCategory{}, errors.New(result.Error.Error())
	}

	return subcategory, nil
}

func SubCategoryDelete(id string) (string, error) {
	var subcategory models.SubCategory

	result := config.DB.Delete(&subcategory, id)

	if result.Error != nil {
		return "", errors.New(result.Error.Error())
	}

	return "subcategory deleted successfully", nil
}
