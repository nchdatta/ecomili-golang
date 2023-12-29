package validations

import (
	_ "github.com/go-playground/validator/v10"
)

type CategoryCreate struct {
	Name      string   `json:"name" validate:"required,min=3,max=150"`
	Icon      string   `json:"icon"`
	Tags      []string `json:"tags" validate:"required"`
	IsSpecial bool     `json:"is_special"`
}
type CategoryUpdate struct {
	Name      string   `json:"name" validate:"required,min=3,max=150"`
	Icon      string   `json:"icon"`
	Tags      []string `json:"tags"`
	IsSpecial bool     `json:"is_special"`
}
