package services

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
	"gorm.io/gorm"
)

func GetAllCategories() (*[]models.Category, error) {
	categories := []models.Category{}

	if err := database.DBConn.Preload("Tags").Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}
func GetCategoryByID(id int) (*models.Category, error) {
	category := &models.Category{}

	if errFound := database.DBConn.Where("id = ?", id).First(&category).Error; errFound != nil {
		if errFound == gorm.ErrRecordNotFound {
			return nil, errors.New("category not found")
		}
		return nil, errFound
	}
	return category, nil
}
func CreateCategory(categoryCreate *validations.CategoryCreate) (*models.Category, error) {
	var existingCategory models.Category
	existErr := database.DBConn.Where("name = ?", strings.ToLower(categoryCreate.Name)).Find(&existingCategory).Error

	if existingCategory.Name != "" {
		return nil, errors.New("Category already exists with the name: " + categoryCreate.Name)
	}
	if existErr != nil {
		return nil, existErr
	}

	category := &models.Category{
		Name:      categoryCreate.Name,
		Icon:      sql.NullString{String: categoryCreate.Icon, Valid: categoryCreate.Icon != ""},
		IsSpecial: categoryCreate.IsSpecial,
	}

	if err := database.DBConn.Create(&category).Error; err != nil {
		return nil, err
	}

	if len(categoryCreate.Tags) > 0 {
		for i := range categoryCreate.Tags {
			tag := &models.CategoryTag{
				Name:       categoryCreate.Tags[i],
				CategoryID: category.ID,
			}
			if err := database.DBConn.Create(tag).Error; err != nil {
				return nil, err
			}
		}
	}

	return category, nil
}

func UpdateCategory(id int, categoryUpdate *validations.CategoryUpdate) (*models.Category, error) {
	category := &models.Category{}
	if err := database.DBConn.Where("id = ?", id).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	category.Name = categoryUpdate.Name
	category.Icon = sql.NullString{String: categoryUpdate.Icon, Valid: categoryUpdate.Icon != ""}
	category.IsSpecial = categoryUpdate.IsSpecial

	for i := range categoryUpdate.Tags {
		tagName := categoryUpdate.Tags[i]
		var existingTag models.CategoryTag
		existErr := database.DBConn.First(&existingTag, "name = ? AND category_id = ?", tagName, category.ID).Error

		if existErr == nil {
			existingTag.Name = tagName
			if err := database.DBConn.Save(&existingTag).Error; err != nil {
				return nil, err
			}
		}

		// Create new tag if it doesn't exist
		if errors.Is(existErr, gorm.ErrRecordNotFound) {
			tag := &models.CategoryTag{
				Name:       tagName,
				CategoryID: category.ID,
			}
			if err := database.DBConn.Create(tag).Error; err != nil {
				return nil, err
			}
		}
	}

	if err := database.DBConn.Save(&category).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteCategory(id int) (*models.Category, error) {
	category := &models.Category{}

	if err := database.DBConn.Where("id = ?", id).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	if err := database.DBConn.Where("category_id = ?", id).Delete(&models.CategoryTag{}).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Delete(&category).Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
