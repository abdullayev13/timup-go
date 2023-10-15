package repo

import (
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
)

type SmsCode struct {
	DB *gorm.DB
}

func (r *SmsCode) Create(model *models.SmsCode) (*models.SmsCode, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *SmsCode) LastByPhoneNumber(phoneNumber string) (*models.SmsCode, error) {
	model := new(models.SmsCode)
	err := r.DB.Where("phone_number = ?", phoneNumber).Order("sent_at DESC").
		First(model).Error

	return model, err
}

func (r *SmsCode) Update(model *models.SmsCode) (*models.SmsCode, error) {
	err := r.DB.Save(model).Error

	return model, err
}
