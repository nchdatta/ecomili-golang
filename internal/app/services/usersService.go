package services

import (
	"errors"

	"github.com/nchdatta/ecomili-golang/internal/app/validations"
	"github.com/nchdatta/ecomili-golang/internal/database"
	"github.com/nchdatta/ecomili-golang/internal/models"
	"gorm.io/gorm"
)

func GetAllUsers() (*[]models.User, error) {
	users := []models.User{}

	if err := database.DBConn.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}
func GetUserByID(id int) (*models.User, error) {
	user := &models.User{}

	result := database.DBConn.Where("id = ?", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("USER NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func CreateUser(userCreate *validations.UserCreate) (*models.User, error) {
	user := &models.User{}

	if result := database.DBConn.Find(&user).Where("email = ?", userCreate.Email); result != nil {
		return nil, errors.New("USER ALREADY EXISTS WITH THE EMAIL")
	}

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
	return user, nil
}

func UpdatedUser(id int, userUpdate *validations.UserUpdate) (*models.User, error) {
	user := &models.User{}

	result := database.DBConn.Where("id = ?", id).First(&user)
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
	return user, nil
}

func DeleteUser(id int) (*models.User, error) {
	user := &models.User{}

	result := database.DBConn.Where("id = ?", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("USER NOT FOUND")
	} else if result.Error != nil {
		return nil, result.Error
	}

	if err := database.DBConn.Delete(&user).Where("id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
