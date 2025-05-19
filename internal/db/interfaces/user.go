package interfaces

import (
	"github.com/Aman5681/notify/internal/db/models"
	"github.com/Aman5681/notify/internal/db/repositories"
)

type UserRepositoryInterface interface {
	InsertUser(user *models.User) error
	GetUserById(userId string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(userId string) error
	DeleteUser(userId string) error
	UpdatePassword(userId string) error
}

var _ UserRepositoryInterface = (*repositories.UserRepository)(nil)
