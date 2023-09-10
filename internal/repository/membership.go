package repository

import "gorm.io/gorm"

type MembershipRepository interface {
}

type membershipRepository struct {
	db *gorm.DB
}

func newMembershipRepository(db *gorm.DB) MembershipRepository {
	return &membershipRepository{db}
}
