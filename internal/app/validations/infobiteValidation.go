package validations

import (
	_ "github.com/go-playground/validator/v10"
)

type InfobiteCreate struct {
	Title   string `json:"title" validate:"required,min=3,max=150"`
	Picture string `json:"picture"`
}
type InfobiteUpdate struct {
	Title   string `json:"title" validate:"required,min=3,max=150"`
	Picture string `json:"picture"`
}
