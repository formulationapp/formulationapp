package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string

	Workspaces []Workspace `gorm:"many2many:memberships"`
}
