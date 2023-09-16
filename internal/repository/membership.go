package repository

import (
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(membership model.Membership) (model.Membership, error)
	GetByUserID(userID uint) ([]model.Membership, error)
	GetByUserIDAndWorkspaceID(userID, workspaceID uint) (model.Membership, error)
}

type membershipRepository struct {
	db *gorm.DB
}

func newMembershipRepository(db *gorm.DB) MembershipRepository {
	err := db.AutoMigrate(&model.Membership{})
	if err != nil {
		logrus.Fatalf("failed to migrate membership model: %s", err)
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
	tx := m.db.Where("user_id = ?", userID).Find(&memberships)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return memberships, nil
}

func (m membershipRepository) GetByUserIDAndWorkspaceID(userID, workspaceID uint) (model.Membership, error) {
	var membership model.Membership
	tx := m.db.Where("user_id = ? and workspace_id = ?", userID, workspaceID).First(&membership)
	if tx.Error != nil {
		return membership, tx.Error
	}
	return membership, nil
}
