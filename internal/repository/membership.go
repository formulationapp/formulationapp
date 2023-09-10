package repository

import (
	"github.com/charmbracelet/log"
	"github.com/formulationapp/formulationapp/internal/model"
	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(membership model.Membership) (model.Membership, error)
	GetByUserID(userID uint) ([]model.Membership, error)
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

func (m membershipRepository) Create(membership model.Membership) (model.Membership, error) {
	err := m.db.Create(&membership).Error
	if err != nil {
		return model.Membership{}, err
	}
	return membership, nil
}

func (m membershipRepository) GetByUserID(userID uint) ([]model.Membership, error) {
	var memberships []model.Membership
	m.db.Where("user_id = ?", userID).Find(&memberships)
	return memberships, nil
}
