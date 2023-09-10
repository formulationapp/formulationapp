package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	FormID    uint
	Field     string
	Value     string
	Submitted bool
}
