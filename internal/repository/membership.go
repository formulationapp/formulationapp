package repository

import (
	"github.com/charmbracelet/log"
	"github.com/formulationapp/formulationapp/internal/model"
	"gorm.io/gorm"
)

type MembershipRepository interface {
}

type membershipRepository struct {
	db *gorm.DB
}

func newMembershipRepository(db *gorm.DB) MembershipRepository {
	err := db.AutoMigrate(&model.Membership{})
	if err != nil {
		log.Fatalf("failed to migrate membership model: %s", err)
	}
	return &membershipRepository{db}
}
