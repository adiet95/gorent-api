package models

import (
	"time"
)

//Struct ke data sesuai dengan data di database

type User struct {
	UserId      string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id,omitempty"`
	UserName    string    `json:"user_name,omitempty"`
	Email       string    `json:"email,omitempty" validate:"required"`
	Role        string    `json:"role,omitempty"`
	Password    string    `json:"Password,omitempty" validate:"required"`
	Path        string    `json:"path,omitempty"`
	FileName    string    `json:"file_name,omitempty"`
	Address     string    `json:"address,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	CreatedAt   time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Users []User
