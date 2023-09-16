package repository

import (
	"fmt"

	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user model.User) (model.User, error)
	FindByID(id uint) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) Save(user model.User) (model.User, error) {
	u.db.Create(&user)
	return user, nil
}

func (u userRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	u.db.First(&user, id)
	if user.ID == 0 {
		return model.User{}, dto.AppError(fmt.Errorf("user with id %d not found", id))
	}
	return user, nil
}

func (u userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	u.db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return model.User{}, dto.AppError(fmt.Errorf("user with email %s not found", email))
	}
	return user, nil
}

func newUserRepository(db *gorm.DB) UserRepository {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		logrus.Fatalf("failed to migrate user model: %s", err)
	}
	return &userRepository{db}
}
