package services

import (
	"errors"
	"strings"

	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
	"gorm.io/gorm"
)

type UserListResponse struct {
	Users       []models.User `json:"users"`
	TotalFound  int           `json:"total_found"`
	CurrentPage int           `json:"current_page"`
	Limit       int           `json:"offset"`
}

func GetAllUsers(page int, pageSize int) (*UserListResponse, error) {
	users := []models.User{}

	offset := (page - 1) * pageSize

	var totalFound int64
	database.DBConn.Model(&models.User{}).Count(&totalFound)

	if err := database.DBConn.Preload("Role").Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	response := &UserListResponse{
		Users:       users,
		TotalFound:  int(totalFound),
		CurrentPage: page,
		Limit:       pageSize,
	}

	return response, nil
}
func GetUserByID(id int) (*models.User, error) {
	user := &models.User{}

	result := database.DBConn.Where("id = ?", id).Preload("Role").First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("USER NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func CreateUser(userCreate *validations.UserCreate) (*models.User, error) {
	var existingUser models.User
	existErr := database.DBConn.Where("email = ?", strings.ToLower(userCreate.Email)).Find(&existingUser).Error

	if existingUser.Email != "" {
		return nil, errors.New("User already exists with the email: " + userCreate.Email)
	}
	if existErr != nil {
		return nil, existErr
	}

	user := &models.User{}
	user.Name = userCreate.Name
	user.Phone = userCreate.Phone
	user.Password = userCreate.Password
	user.Email = userCreate.Email
	user.RoleID = uint(userCreate.RoleID)
	if userCreate.Avatar.Valid {
		user.Avatar = userCreate.Avatar
	}

	if err := database.DBConn.Create(&user).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdatedUser(id int, userUpdate *validations.UserUpdate) (*models.User, error) {
	user := &models.User{}

	result := database.DBConn.Select("id").Where("id = ?", id).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("USER NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	user.Name = userUpdate.Name
	user.Email = userUpdate.Email
	user.Phone = userUpdate.Phone
	user.Password = userUpdate.Password
	user.RoleID = uint(userUpdate.RoleID)

	if userUpdate.Avatar.Valid {
		user.Avatar = userUpdate.Avatar
	}

	if err := database.DBConn.Save(&user).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteUser(id int) (*models.User, error) {
	user := &models.User{}

	result := database.DBConn.Where("id = ?", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("USER NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	if err := database.DBConn.Delete(&user).Unscoped().Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
