package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nchdatta/ecomili-golang/config"
	"github.com/nchdatta/ecomili-golang/internal/app/validations"
)

func AuthLogin(authLogin *validations.LoginCredentials) (string, error) {
	// Token Secret Key
	config, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	// Check if user exists
	user, err := GetUserByEmail(authLogin.Email)
	if err != nil {
		return "", err
	}

	// if !helpers.CheckPasswordHash(authLogin.Password, authLogin.Password) {
	// 	return "", errors.New("invalid password")
	// }

	// Create the Claims
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
