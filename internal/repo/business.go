package repo

import (
	"abdullayev13/timeup/internal/dtos"
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

func (r *Business) GetByCategory(data *dtos.BusinessFilter) ([]*dtos.BusinessMini, error) {
	listModel := make([]*dtos.BusinessMini, 0)

	err := r.DB.Raw(`SELECT b.id as business_id, b.user_id, b.experience,
u.fist_name, u.last_name, u.photo_url FROM users u 
JOIN business_profiles b on b.user_id=u.id 
WHERE b.work_category_id = ? LIMIT ? OFFSET ?`, data.CategoryId, data.Limit, data.Offset).Find(&listModel).Error

	if err != nil {
		return nil, err
	}

	return listModel, nil
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
