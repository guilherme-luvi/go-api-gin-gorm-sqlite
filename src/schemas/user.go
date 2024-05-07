package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}
