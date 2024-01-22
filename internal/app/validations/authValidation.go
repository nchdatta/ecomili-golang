package validations

import (
	_ "github.com/go-playground/validator/v10"
)

type LoginCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
