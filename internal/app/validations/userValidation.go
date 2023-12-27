package validations

import (
	"database/sql"

	_ "github.com/go-playground/validator/v10"
)

type UserCreate struct {
	Name     string         `json:"name" validate:"required,max=100"`
	Phone    string         `json:"phone" validate:"max:12"`
	Email    string         `json:"email" validate:"required,email,max:150"`
	Password string         `json:"password" validate:"required,min:6,max:200"`
	Avatar   sql.NullString `json:"avatar"`
	RoleID   int            `json:"role_id"`
}
type UserUpdate struct {
	Name     string         `json:"name" validate:"required,max=100"`
	Phone    string         `json:"phone" validate:"max:12"`
	Email    string         `json:"email" validate:"required,email,max:150"`
	Password string         `json:"password" validate:"min:6,max:200"`
	Avatar   sql.NullString `json:"avatar"`
	RoleID   int            `json:"role_id"`
}
