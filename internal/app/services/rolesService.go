package services

import (
	"errors"
	"strings"

	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
	"gorm.io/gorm"
)

type RoleListResponse struct {
	Roles       []models.Role `json:"roles"`
	Pages       int           `json:"pages"`
	CurrentPage int           `json:"current_page"`
	Limit       int           `json:"offset"`
}

func GetAllRoles(page int, pageSize int) (*RoleListResponse, error) {
	roles := []models.Role{}

	offset := (page - 1) * pageSize
	var totalFound int64
	database.DBConn.Model(&models.Role{}).Count(&totalFound)

	if err := database.DBConn.Offset(offset).Limit(pageSize).Find(&roles).Error; err != nil {
		return nil, err
	}

	response := &RoleListResponse{
		Roles:       roles,
		Pages:       int(totalFound),
		CurrentPage: page,
		Limit:       pageSize,
	}

	return response, nil
}
func GetRoleByID(id int) (*models.Role, error) {
	role := &models.Role{}

	result := database.DBConn.Where("id = ?", id).First(&role)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ROLE NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}
func CreateRole(roleCreate *validations.RoleCreate) (*models.Role, error) {
	var existingRole models.Role
	existErr := database.DBConn.Select("name").Where("name = ?", strings.ToLower(roleCreate.Name)).Find(&existingRole).Error

	if existingRole.Name != "" {
		return nil, errors.New("Role already exists with the name: " + roleCreate.Name)
	}
	if existErr != nil {
		return nil, existErr
	}

	role := &models.Role{
		Name: roleCreate.Name,
	}
	if err := database.DBConn.Create(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func UpdatedRole(id int, roleUpdate *validations.RoleUpdate) (*models.Role, error) {
	var existingRole models.Role

	result := database.DBConn.Select("id").Where("id = ?", id).First(&existingRole)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ROLE NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	role := &models.Role{
		Name: roleUpdate.Name,
	}

	if err := database.DBConn.Save(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func DeleteRole(id int) (*models.Role, error) {
	role := &models.Role{}

	result := database.DBConn.Where("id = ?", id).First(&role)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ROLE NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	if err := database.DBConn.Where("id = ?", id).Unscoped().Delete(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
