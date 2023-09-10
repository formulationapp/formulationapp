package repository

import "gorm.io/gorm"

type AnswerRepository interface {
}

type answerRepository struct {
	db *gorm.DB
}

func newAnswerRepository(db *gorm.DB) AnswerRepository {
	return &answerRepository{db}
}
