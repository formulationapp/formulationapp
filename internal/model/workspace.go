package model

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	Name string
}
