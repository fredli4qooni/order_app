package repository

import (
	"gorm.io/gorm"
	"it-order-app/backend/model"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(id uint) (model.User, error) // Menerima uint
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id uint) (model.User, error) { // Menerima uint
	var user model.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
