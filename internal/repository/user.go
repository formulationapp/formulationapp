package repository

import (
	"github.com/charmbracelet/log"
	"github.com/formulationapp/formulationapp/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
}

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) UserRepository {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("failed to migrate user model: %s", err)
	}
	return &userRepository{db}
}
