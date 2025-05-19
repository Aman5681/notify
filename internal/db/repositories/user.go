package repositories

import (
	"github.com/Aman5681/notify/internal/db/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserById(userId string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepository) GetUserByEmail(emailId string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepository) UpdateUser(userId string) error {
	return nil
}

func (r *UserRepository) DeleteUser(userId string) error {
	return nil
}

func (r *UserRepository) UpdatePassword(userId string) error {
	return nil
}
