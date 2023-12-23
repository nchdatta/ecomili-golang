package services

import (
	"github.com/google/uuid"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
)

func GetAllUsers() ([]models.Role, error) {
	roles := []models.Role{}

	if err := database.DBConn.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}
func GetUserByID(id string) (*models.Role, error) {
	role := &models.Role{}

	if err := database.DBConn.Find(&role).Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return role, nil
}
func CreateUser(roleCreate *validations.UserCreate) (*models.Role, error) {
	role := models.Role{
		ID:   uuid.New(),
		Name: roleCreate.Name,
	}

	if err := database.DBConn.Find(&role).Where("name=?", roleCreate.Name).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Create(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func UpdatedUser(id string, roleUpdate *validations.UserUpdate) (*models.Role, error) {
	role := &models.Role{
		ID:   uuid.New(),
		Name: roleUpdate.Name,
	}

	if err := database.DBConn.Find(&role).Where("id=?", id).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Find(&role).Where("name=?", roleUpdate.Name).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Save(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func DeleteUser(id string) (*models.Role, error) {
	role := &models.Role{}

	if err := database.DBConn.Find(&role).Where("id=?", id).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Delete(&role).Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return role, nil
}
