package validations

import (
	_ "github.com/go-playground/validator/v10"
)

type AuthorCreate struct {
	Name        string `json:"name" validate:"required"`
	Designation string `json:"designation"`
	Bio         string `json:"bio"`
	Avatar      string `json:"avatar"`
	Twitter     string `json:"twitter"`
	Linkedin    string `json:"linkedin"`
	Email       string `json:"email"`
}

type AuthorUpdate struct {
	Name        string `json:"name" validate:"required"`
	Designation string `json:"designation"`
	Bio         string `json:"bio"`
	Avatar      string `json:"avatar"`
	Twitter     string `json:"twitter"`
	Linkedin    string `json:"linkedin"`
	Email       string `json:"email"`
}
