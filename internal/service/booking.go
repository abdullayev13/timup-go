package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/repo"
	"errors"
	"gorm.io/gorm"
)

type Booking struct {
	Repo *repo.Repo
}

func (s *Booking) Create(data *dtos.Booking) (*dtos.Booking, error) {
	data.ID = 0

	_, err := s.Repo.Business.GetById(data.BusinessId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("business not found")
		}
		return nil, err
	}

	model := data.MapToModel()

	model, err = s.Repo.Booking.Create(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)

	return data, nil
}

func (s *Booking) GetList(data *dtos.BookingFilter) ([]*dtos.Booking, error) {
	return s.getList(data)
}

func (s *Booking) getList(data *dtos.BookingFilter) ([]*dtos.Booking, error) {
	if data.Limit == 0 {
		data.Limit = 100
	}

	list, err := s.Repo.Booking.GetList(data)
	if err != nil {
		return nil, err
	}

	listDto := make([]*dtos.Booking, len(list))
	for i, model := range list {
		listDto[i] = new(dtos.Booking).
			MapFromModel(model)
	}

	return listDto, nil
}

func (s *Booking) GetListByClient(data *dtos.BookingFilter, userId int) ([]*dtos.Booking, error) {
	data.BusinessId = 0
	data.ClientId = userId

	return s.getList(data)
}

func (s *Booking) GetListByBusiness(data *dtos.BookingFilter) ([]*dtos.Booking, error) {
	data.ClientId = 0

	return s.getList(data)
}

func (s *Booking) DeleteById(id int) error {
	return s.Repo.Booking.DeleteById(id)
}

func (s *Booking) DeleteByClient(id, userId int) error {
	booking, err := s.Repo.Booking.GetById(id)
	if err != nil {
		return err
	}

	user, err := s.Repo.Users.GetById(userId)
	if err != nil || user.ID != booking.ClientId {
		return errors.New("access denied")
	}

	return s.Repo.Booking.DeleteById(id)
}

func (s *Booking) DeleteByBusiness(id, userId int) error {
	booking, err := s.Repo.Booking.GetById(id)
	if err != nil {
		return err
	}

	business, err := s.Repo.Business.GetByUserId(userId)
	if err != nil && business.ID != booking.BusinessId {
		return errors.New("access denied")
	}

	return s.Repo.Booking.DeleteById(id)
}
