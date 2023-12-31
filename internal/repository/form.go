package repository

import (
	"encoding/json"
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FormRepository interface {
	Create(form model.Form) (model.Form, error)
	Update(form model.Form) (model.Form, error)
	GetByWorkspaceID(workspaceID uint) ([]model.Form, error)
	GetByIDAndWorkspaceID(id, workspaceID uint) (model.Form, error)
	Delete(form model.Form) error
	GetBySecret(secret string) (model.Form, error)
}

type formRepository struct {
	db *gorm.DB
}

func (f formRepository) GetBySecret(secret string) (model.Form, error) {
	var form model.Form
	tx := f.db.Where("secret = ?", secret).First(&form)
	if tx.Error != nil {
		return form, tx.Error
	}

	err := json.Unmarshal([]byte(form.Definition), &form.Data)
	if err != nil {
		return form, err
	}

	return form, nil
}

func (f formRepository) GetByIDAndWorkspaceID(id, workspaceID uint) (model.Form, error) {
	var form model.Form
	tx := f.db.Where("id = ? and workspace_id = ?", id, workspaceID).First(&form)
	if tx.Error != nil {
		return form, tx.Error
	}

	err := json.Unmarshal([]byte(form.Definition), &form.Data)
	if err != nil {
		return form, err
	}

	return form, nil
}

func (f formRepository) Create(form model.Form) (model.Form, error) {
	j, err := json.Marshal(form.Data)
	if err != nil {
		return model.Form{}, err
	}
	form.Definition = string(j)

	tx := f.db.Create(&form)
	if tx.Error != nil {
		return form, tx.Error
	}
	return form, nil
}

func (f formRepository) Update(form model.Form) (model.Form, error) {
	j, err := json.Marshal(form.Data)
	if err != nil {
		return model.Form{}, err
	}
	form.Definition = string(j)

	tx := f.db.Save(form)
	if tx.Error != nil {
		return form, tx.Error
	}
	return form, nil
}

func (f formRepository) GetByWorkspaceID(workspaceID uint) ([]model.Form, error) {
	var forms []model.Form
	tx := f.db.Where("workspace_id = ?", workspaceID).Find(&forms)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var filledForms []model.Form
	for _, form := range forms {
		err := json.Unmarshal([]byte(form.Definition), &form.Data)
		if err != nil {
			return nil, err
		}
		filledForms = append(filledForms, form)
	}

	return filledForms, nil
}

func (f formRepository) Delete(form model.Form) error {
	return f.db.Delete(&form).Error
}

func newFormRepository(db *gorm.DB) FormRepository {
	err := db.AutoMigrate(model.Form{})
	if err != nil {
		logrus.Fatalf("failed to migrate form model: %s", err)
	}
	return &formRepository{db}
}
