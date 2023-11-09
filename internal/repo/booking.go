package repo

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
	"strings"
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

// other

func (r *Booking) GetListByClient(data *dtos.BookingFilter) ([]*dtos.BookingMini, error) {
	query := []string{`SELECT b.*,
       u.fist_name,
       u.last_name,
       u.user_name,
       u.phone_number,
       u.photo_url
FROM bookings b
         JOIN business_profiles bp ON bp.id = b.business_id
         JOIN users u ON u.id = bp.user_id
WHERE b.client_id = ?`}
	args := []any{data.ClientId}

	if data.BusinessId != 0 {
		query = append(query, "AND b.business_id = ?")
		args = append(args, data.BusinessId)
	}
	if data.Coming {
		query = append(query, "AND b.date > now()")
	}
	if data.Date != "" {
		query = append(query, "AND b.date::date = to_date(?, 'DD/MM/YYYY')")
		args = append(args, data.Date)
	}
	query = append(query, "ORDER BY b.date LIMIT ? OFFSET ?")
	args = append(args, data.Limit, data.Offset)

	res := make([]*dtos.BookingMini, 0, data.Limit)
	err := r.DB.Raw(strings.Join(query, " "), args...).Find(&res).Error

	return res, err
}

func (r *Booking) GetListByBusiness(data *dtos.BookingFilter) ([]*dtos.BookingMini, error) {
	query := []string{`SELECT b.*,
       u.fist_name,
       u.last_name,
       u.user_name,
       u.phone_number,
       u.photo_url
FROM bookings b
         JOIN users u ON u.id = b.client_id
WHERE b.business_id = ?`}
	args := []any{data.BusinessId}

	if data.ClientId != 0 {
		query = append(query, "AND b.client_id = ?")
		args = append(args, data.ClientId)
	}
	if data.Coming {
		query = append(query, "AND b.date > now()")
	}
	if data.Date != "" {
		query = append(query, "AND b.date::date = to_date(?, 'DD/MM/YYYY')")
		args = append(args, data.Date)
	}
	query = append(query, "ORDER BY b.date LIMIT ? OFFSET ?")
	args = append(args, data.Limit, data.Offset)

	res := make([]*dtos.BookingMini, 0, data.Limit)
	err := r.DB.Raw(strings.Join(query, " "), args...).Find(&res).Error

	return res, err
}
