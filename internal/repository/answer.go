package repository

import (
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AnswerRepository interface {
	Create(answer model.Answer) (model.Answer, error)
	GetByFormID(formID uint) ([]model.Answer, error)
}

type answerRepository struct {
	db *gorm.DB
}

func (a answerRepository) Create(answer model.Answer) (model.Answer, error) {
	tx := a.db.Create(&answer)
	if tx.Error != nil {
		return model.Answer{}, tx.Error
	}
	return answer, nil
}

func (a answerRepository) GetByFormID(formID uint) ([]model.Answer, error) {
	var answers []model.Answer
	tx := a.db.Where("form_id = ?", formID).Find(&answers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return answers, nil
}

func newAnswerRepository(db *gorm.DB) AnswerRepository {
	err := db.AutoMigrate(model.Answer{})
	if err != nil {
		logrus.Fatalf("failed to migrate answer model: %s", err)
	}
	return &answerRepository{db}
}
