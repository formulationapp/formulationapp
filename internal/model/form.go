package model

import "gorm.io/gorm"

type Form struct {
	gorm.Model
	Name        string
	Definition  string `gorm:"type:text"`
	WorkspaceID uint
}
