package repository

import (
	"github.com/charmbracelet/log"
	"github.com/formulationapp/formulationapp/internal/model"
	"gorm.io/gorm"
)

type WorkspaceRepository interface {
}

type workspaceRepository struct {
	db *gorm.DB
}

func newWorkspaceRepository(db *gorm.DB) WorkspaceRepository {
	err := db.AutoMigrate(model.Workspace{})
	if err != nil {
		log.Fatalf("failed to migrate workspace model: %s", err)
	}
	return &workspaceRepository{db}
}
