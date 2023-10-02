package repo

import (
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
)

type Category struct {
	DB *gorm.DB
}

func (r *Category) Create(model *models.WorkCategory) (*models.WorkCategory, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *Category) GetByParentId(parentId int) ([]*models.WorkCategory, error) {
	slc := make([]*models.WorkCategory, 0, 10)
	err := r.DB.Where("parent_id = ?", parentId).
		Find(&slc).Error
	if err != nil {
		return nil, err
	}

	return slc, nil
}

func (r *Category) GetById(id int) (*models.WorkCategory, error) {
	model := new(models.WorkCategory)
	err := r.DB.First(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Category) Update(model *models.WorkCategory) (*models.WorkCategory, error) {
	err := r.DB.Save(model).Error

	return model, err
}

func (r *Category) DeleteById(id int) error {
	model := new(models.WorkCategory)
	err := r.DB.Where("id = ?", id).Delete(model).Error

	return err
}

func (r *Category) DeleteByParentId(parentId int) error {
	model := new(models.WorkCategory)
	err := r.DB.Where("parent_id = ?", parentId).Delete(model).Error

	return err
}
