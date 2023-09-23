package model

import "gorm.io/gorm"

type Form struct {
	gorm.Model
	Name        string `json:"name"`
	Definition  string `gorm:"type:text" json:"definition"`
	WorkspaceID uint   `json:"workspaceID"`
	Secret      string `json:"secret"`

	Data map[string]interface{} `gorm:"-" json:"data"`
}
