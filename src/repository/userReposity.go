package repository

import (
	"e-wallet/src/models"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	FindById(id int) ([]*models.User, error)
	FindByName(name string) ([]*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}


type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}


func (r *userRepository) FindAll() ([]*models.User, error) {
	var users []*models.User

	err := r.db.Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
