package repository

import "gorm.io/gorm"

type WorkspaceRepository interface {
}

type workspaceRepository struct {
	db *gorm.DB
}

func newWorkspaceRepository(db *gorm.DB) WorkspaceRepository {
	return &workspaceRepository{db}
}
