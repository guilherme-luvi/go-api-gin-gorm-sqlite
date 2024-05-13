package repositories

import (
	"github.com/guilherme-luvi/go-api-gin-swagger-goorm-sqlite/src/schemas"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{db}
}

func (u *user) CreateUser(user schemas.User) error {
	return u.db.Create(&user).Error
}

func (u *user) GetUserById(id string) (schemas.User, error) {
	user := schemas.User{}
	err := u.db.Where("id = ?", id).First(&user).Error
	return user, err
}
