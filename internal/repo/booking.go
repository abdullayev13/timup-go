package repo

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
)

type Booking struct {
	DB *gorm.DB
}

func (r *Booking) Create(model *models.Booking) (*models.Booking, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *Booking) GetById(id int) (*models.Booking, error) {
	model := new(models.Booking)
	err := r.DB.First(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Booking) GetList(data *dtos.BookingFilter) ([]*models.Booking, error) {
	tx := r.DB.Order("date")

	if data.Offset != 0 {
		tx.Offset(data.Offset)
	}
	if data.Limit != 0 {
		tx.Limit(data.Limit)
	}
	if data.BusinessId != 0 {
		tx.Where("business_id = ?", data.BusinessId)
	}
	if data.ClientId != 0 {
		tx.Where("client_id = ?", data.ClientId)
	}
	if data.Coming {
		tx.Where("date > now()")
	}
	if data.Date != "" {
		tx.Where("date::date = to_date(?, 'DD/MM/YYYY')", data.Date)
	}

	models := make([]*models.Booking, 0)

	err := tx.Find(&models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (r *Booking) Update(model *models.Booking) (*models.Booking, error) {
	err := r.DB.Save(model).Error

	return model, err
}

func (r *Booking) DeleteById(id int) error {
	model := new(models.Booking)
	err := r.DB.Where("id = ?", id).Delete(model).Error

	return err
}
