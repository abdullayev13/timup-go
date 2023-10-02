package repo

import (
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
)

type Business struct {
	DB *gorm.DB
}

func (r *Business) Create(model *models.BusinessProfile) (*models.BusinessProfile, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *Business) GetByUserId(userId int) (*models.BusinessProfile, error) {
	model := new(models.BusinessProfile)
	err := r.DB.Where("user_id = ?", userId).
		First(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Business) GetById(id int) (*models.BusinessProfile, error) {
	model := new(models.BusinessProfile)
	err := r.DB.First(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Business) Update(model *models.BusinessProfile) (*models.BusinessProfile, error) {
	err := r.DB.Save(model).Error

	return model, err
}

func (r *Business) DeleteById(id int) error {
	model := new(models.BusinessProfile)
	err := r.DB.Where("id = ?", id).Delete(model).Error

	return err
}

func (r *Business) DeleteByUserId(userId int) error {
	model := new(models.BusinessProfile)
	err := r.DB.Where("user_id = ?", userId).Delete(model).Error

	return err
}
