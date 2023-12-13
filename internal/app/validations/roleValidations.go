package validations

import (
	_ "github.com/go-playground/validator/v10"
)

type RoleCreate struct {
	Name string `json:"name" validate:"required,max=100"`
}
type RoleUpdate struct {
	Name string `json:"name" validate:"required,max=100"`
}
