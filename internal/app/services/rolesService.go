package services

import (
	"errors"
	"strings"

	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
	"gorm.io/gorm"
)

func GetAllRoles() ([]models.Role, error) {
	roles := []models.Role{}

	if err := database.DBConn.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
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
