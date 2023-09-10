package model

import (
	"gorm.io/gorm"
	"time"
)

type Membership struct {
	UserID      uint `gorm:"primaryKey"`
	WorkspaceID uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
