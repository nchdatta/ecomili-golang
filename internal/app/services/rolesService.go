package services

import (
	"errors"

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
	role := models.Role{
		Name: roleCreate.Name,
	}

	result := database.DBConn.Where("name = ?", roleCreate.Name).First(&role)
	if result != nil {
		return nil, errors.New("ROLE ALREADY EXISTS")
	}

	if err := database.DBConn.Create(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func UpdatedRole(id int, roleUpdate *validations.RoleUpdate) (*models.Role, error) {
	role := &models.Role{}

	result := database.DBConn.Where("id = ?", id).First(&role)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ROLE NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	role.Name = roleUpdate.Name
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

	if err := database.DBConn.Where("id = ?", id).Delete(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
