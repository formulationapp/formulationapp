package repository

import (
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WorkspaceRepository interface {
	Create(workspace model.Workspace) (model.Workspace, error)
	GetByID(id uint) (model.Workspace, error)
}

type workspaceRepository struct {
	db *gorm.DB
}

func newWorkspaceRepository(db *gorm.DB) WorkspaceRepository {
	err := db.AutoMigrate(model.Workspace{})
	if err != nil {
		logrus.Fatalf("failed to migrate workspace model: %s", err)
	}
	return &workspaceRepository{db}
}

func (w workspaceRepository) Create(workspace model.Workspace) (model.Workspace, error) {
	err := w.db.Create(&workspace).Error
	if err != nil {
		return model.Workspace{}, err
	}
	return workspace, nil
}

func (w workspaceRepository) GetByID(id uint) (model.Workspace, error) {
	var workspace model.Workspace
	w.db.First(&workspace, id)
	if workspace.ID == 0 {
		return model.Workspace{}, gorm.ErrRecordNotFound
	}
	return workspace, nil
}
