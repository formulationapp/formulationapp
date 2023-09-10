package repository

import (
	"github.com/charmbracelet/log"
	"github.com/formulationapp/formulationapp/internal/model"
	"gorm.io/gorm"
)

type AnswerRepository interface {
}

type answerRepository struct {
	db *gorm.DB
}

func newAnswerRepository(db *gorm.DB) AnswerRepository {
	err := db.AutoMigrate(model.Answer{})
	if err != nil {
		log.Fatalf("failed to migrate answer model: %s", err)
	}
	return &answerRepository{db}
}
