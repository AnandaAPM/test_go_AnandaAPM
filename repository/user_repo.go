package repository

import (
	// "errors"
	"go_test/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUser(username string) (*models.User, error)
}
type userRepository struct {
	DB *gorm.DB
}


func UserRepo(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}
 
func (r *userRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) FindUser(username string) (*models.User, error) {
	var exist models.User

	if err := r.DB.Where("username = ?", username).First(&exist).Error; err != nil {
		return nil, err
	}

	return &exist, nil
}
