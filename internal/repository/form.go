package repository

import "gorm.io/gorm"

type FormRepository interface {
}

type formRepository struct {
	db *gorm.DB
}

func newFormRepository(db *gorm.DB) FormRepository {
	return &formRepository{db}
}
