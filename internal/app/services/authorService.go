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

func GetAllAuthors() ([]models.Author, error) {
	author := []models.Author{}

	if err := database.DBConn.Find(&author).Error; err != nil {
		return nil, err
	}

	return author, nil
}
func GetAuthorByID(id int) (*models.Author, error) {
	author := &models.Author{}

	if err := database.DBConn.Where("id = ?", id).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("author not found")
		}
		return nil, err
	}

	return author, nil
}
func CreateAuthor(authorCreate *validations.AuthorCreate) (*models.Author, error) {
	var existingAuthor models.Role
	existErr := database.DBConn.Select("name").Where("name = ?", strings.ToLower(authorCreate.Name)).Find(&existingAuthor).Error

	if existingAuthor.Name != "" {
		return nil, errors.New("Author already exists with the name: " + authorCreate.Name)
	}
	if existErr != nil {
		return nil, existErr
	}

	author := &models.Author{
		Name: authorCreate.Name,
	}
	if err := database.DBConn.Create(&author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func UpdateAuthor(id int, authorUpdate *validations.AuthorUpdate) (*models.Author, error) {
	author := &models.Author{}

	if err := database.DBConn.Where("id = ?", id).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("author not found")
		}
		return nil, err
	}

	author.Name = authorUpdate.Name
	author.Designation = authorUpdate.Designation
	author.Email = authorUpdate.Email
	author.Bio = authorUpdate.Bio
	author.Avatar = sql.NullString{String: authorUpdate.Avatar, Valid: authorUpdate.Avatar != ""}
	author.Linkedin = authorUpdate.Linkedin
	author.Twitter = authorUpdate.Twitter

	if err := database.DBConn.Save(&author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func DeleteAuthor(id int) (*models.Author, error) {
	author := &models.Author{}

	if err := database.DBConn.Where("id = ?", id).First(&author).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("author not found")
		}
		return nil, err
	}

	if err := database.DBConn.Where("id = ?", id).Delete(&author).Error; err != nil {
		return nil, err
	}
	return author, nil
}
