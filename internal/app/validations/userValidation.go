package validations

import (
	_ "github.com/go-playground/validator/v10"
)

type UserCreate struct {
	Name string `json:"name" validate:"required,max=100"`
}
type UserUpdate struct {
	Name string `json:"name" validate:"required,max=100"`
}
