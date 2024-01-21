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

type InfobiteListResponse struct {
	Infobites   []models.Infobite `json:"infobites"`
	Pages       int               `json:"pages"`
	CurrentPage int               `json:"current_page"`
	Limit       int               `json:"offset"`
}

func GetAllInfobites(page int, pageSize int) (*InfobiteListResponse, error) {
	infobites := []models.Infobite{}

	offset := (page - 1) * pageSize

	var totalFound int64
	database.DBConn.Model(&models.Infobite{}).Count(&totalFound)

	if err := database.DBConn.Offset(offset).Limit(pageSize).Find(&infobites).Error; err != nil {
		return nil, err
	}

	response := &InfobiteListResponse{
		Infobites:   infobites,
		Pages:       int(totalFound),
		CurrentPage: page,
		Limit:       pageSize,
	}

	return response, nil
}
func GetInfobiteByID(id int) (*models.Infobite, error) {
	infobite := &models.Infobite{}

	result := database.DBConn.Where("id = ?", id).First(&infobite)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("infobite NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return infobite, nil
}
func CreateInfobite(infobiteCreate *validations.InfobiteCreate) (*models.Infobite, error) {
	var existingInfobite models.Infobite
	existErr := database.DBConn.Where("title = ?", strings.ToLower(infobiteCreate.Title)).Find(&existingInfobite).Error

	if existingInfobite.Title != "" {
		return nil, errors.New("infobites already exists with the title: " + infobiteCreate.Title)
	}
	if existErr != nil {
		return nil, existErr
	}

	infobite := &models.Infobite{
		Title:   infobiteCreate.Title,
		Picture: sql.NullString{String: infobiteCreate.Picture, Valid: infobiteCreate.Picture != ""},
	}

	if err := database.DBConn.Create(&infobite).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdateInfobite(id int, infobiteUpdate *validations.InfobiteUpdate) (*models.Infobite, error) {
	existingInfobite := &models.Infobite{}

	result := database.DBConn.Select("id").Where("id = ?", id).First(existingInfobite)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("infobite NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	infobite := &models.Infobite{
		Title:   infobiteUpdate.Title,
		Picture: sql.NullString{String: infobiteUpdate.Picture, Valid: infobiteUpdate.Picture != ""},
	}

	if err := database.DBConn.Save(&infobite).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteInfobite(id int) (*models.Infobite, error) {
	infobite := &models.Infobite{}

	result := database.DBConn.Where("id = ?", id).First(&infobite)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("infobite NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	if err := database.DBConn.Delete(&infobite).Unscoped().Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
