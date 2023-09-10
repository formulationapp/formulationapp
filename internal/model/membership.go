package model

import (
	"gorm.io/gorm"
	"time"
)

type Membership struct {
	UserID      int `gorm:"primaryKey"`
	WorkspaceID int `gorm:"primaryKey"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
