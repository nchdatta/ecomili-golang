package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID  `json:"id" gorm:"type:binary(16);default:uuid_generate_v4()"`
	Name      string     `json:"name" gorm:"uniqueIndex;type:varchar(200)"`
	RoleID    *uuid.UUID `json:"role_id" gorm:"type:binary(16)"`
	Role      *Role      `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Role struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:binary(16);defaut:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"uniqueIndex;type:varchar(100)"`
	Status    Status    `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []User    `json:"users"`
}

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)
