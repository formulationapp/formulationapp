package model

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:memberships"`
}
