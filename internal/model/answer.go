package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	FormID     uint
	Submission string
	Field      string
	Value      string
	Submitted  bool
}
