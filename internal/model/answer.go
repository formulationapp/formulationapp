package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	FormID     uint   `json:"formID"`
	Submission string `json:"submission"`
	Field      string `json:"field"`
	Value      string `json:"value"`
	Submitted  bool   `json:"submitted"`
}
