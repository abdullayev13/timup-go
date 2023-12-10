package repo

import (
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
)

type BookingCategory struct {
	DB *gorm.DB
}

func (r *BookingCategory) Create(model *models.BookingCategory) (*models.BookingCategory, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *BookingCategory) GetById(id int) (*models.BookingCategory, error) {
	model := new(models.BookingCategory)
	err := r.DB.First(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *BookingCategory) GetByBusinessId(businessId int) ([]*models.BookingCategory, error) {
	slc := make([]*models.BookingCategory, 0, 10)
	err := r.DB.Where("business_id = ?", businessId).
		Find(&slc).Error
	if err != nil {
		return nil, err
	}

	return slc, nil
}

func (r *BookingCategory) DeleteById(id int) error {
	model := new(models.BookingCategory)
	err := r.DB.Where("id = ?", id).Delete(model).Error

	return err
}
