package repository

import (
	"github.com/charmbracelet/log"
	"github.com/formulationapp/formulationapp/internal/model"
	"gorm.io/gorm"
)

type FormRepository interface {
}

type formRepository struct {
	db *gorm.DB
}

func newFormRepository(db *gorm.DB) FormRepository {
	err := db.AutoMigrate(model.Form{})
	if err != nil {
		log.Fatalf("failed to migrate form model: %s", err)
	}
	return &formRepository{db}
}
