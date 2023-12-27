package services

import (
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
)

func GetAllUsers() (*[]models.User, error) {
	users := []models.User{}

	if err := database.DBConn.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}
func GetUserByID(id string) (*models.User, error) {
	user := &models.User{}

	if err := database.DBConn.Find(&user).Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func CreateUser(userCreate *validations.UserCreate) (*models.User, error) {
	user := &models.User{
		Name:     userCreate.Name,
		Phone:    userCreate.Phone,
		Password: userCreate.Password,
		Email:    userCreate.Email,
		RoleID:   uint(userCreate.RoleID),
	}

	if userCreate.Avatar.Valid {
		user.Avatar = userCreate.Avatar
	}

	if err := database.DBConn.Find(&user).Where("name=?", userCreate.Name).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdatedUser(id string, userUpdate *validations.UserUpdate) (*models.User, error) {
	user := &models.User{
		Name:     userUpdate.Name,
		Email:    userUpdate.Email,
		Phone:    userUpdate.Phone,
		Password: userUpdate.Password,
		RoleID:   uint(userUpdate.RoleID),
	}

	if userUpdate.Avatar.Valid {
		user.Avatar = userUpdate.Avatar
	}

	if err := database.DBConn.Find(&user).Where("id=?", id).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Find(&user).Where("name=?", userUpdate.Name).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id string) (*models.User, error) {
	user := &models.User{}

	if err := database.DBConn.Find(&user).Where("id=?", id).Error; err != nil {
		return nil, err
	}

	if err := database.DBConn.Delete(&user).Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
