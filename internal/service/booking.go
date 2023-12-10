package service

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
	"errors"
	"fmt"
	"time"
)

type Booking struct {
	Repo *repo.Repo
}

func (s *Booking) Create(data *dtos.Booking) (*dtos.Booking, error) {
	data.ID = 0
	var err error

	model := data.MapToModel()

	if model.BookingCategoryId == nil {
		exists := s.Repo.Business.ExistsById(model.BusinessId)
		if !exists {
			return nil, errors.New("business not found")
		}
		model.EndTime = model.Date.Add(config.DefaultBookingDuration)
	} else {
		bc, err := s.Repo.BookingCategory.GetById(*model.BookingCategoryId)
		if err != nil {
			return nil, errors.New("booking_category: " + err.Error())
		}

		model.BusinessId = bc.BusinessId
		model.EndTime = model.Date.Add(bc.Duration)
	}

	if model.Date.Before(time.Now()) {
		return nil, errors.New("booking time is past or not given")
	}

	bookings := s.Repo.Booking.GetBetweenByBusiness(model.BusinessId, model.Date, model.EndTime)
	if len(bookings) > 0 {
		return nil, fmt.Errorf("can not create: exists booking from '%s' to '%s'",
			utill.FormatHHmmTZ0(bookings[0].Date),
			utill.FormatHHmmTZ0(bookings[0].EndTime))
	}

	model, err = s.Repo.Booking.Create(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)

	return data, nil
}

func (s *Booking) GetList(data *dtos.BookingFilter) ([]*dtos.Booking, error) {
	list, err := s.getList(data)
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

func (s *Booking) getList(data *dtos.BookingFilter) ([]*models.Booking, error) {
	if data.Limit == 0 {
		data.Limit = 100
	}

	return s.Repo.Booking.GetList(data)
	//if err != nil {
	//	return nil, err
	//}
	//
	//listDto := make([]*dtos.Booking, len(list))
	//for i, model := range list {
	//	listDto[i] = new(dtos.Booking).
	//		MapFromModel(model)
	//}
	//
	//return listDto, nil
}

func (s *Booking) GetListByClient(data *dtos.BookingFilter, userId int) ([]*dtos.BookingMini, error) {
	data.ClientId = userId
	if data.Limit == 0 {
		data.Limit = 100
	}

	list, err := s.Repo.Booking.GetListByClient(data)
	if err != nil {
		return nil, err
	}

	for _, bus := range list {
		bus.DateJson, bus.Time = utill.Format(bus.Date)
		bus.PhotoUrl = utill.PutMediaDomain(bus.PhotoUrl)
	}

	return list, nil
}

func (s *Booking) GetListByBusiness(data *dtos.BookingFilter, businessId int) ([]*dtos.BookingMini, error) {
	data.BusinessId = businessId
	if data.Limit == 0 {
		data.Limit = 100
	}

	list, err := s.Repo.Booking.GetListByBusiness(data)
	if err != nil {
		return nil, err
	}

	for _, bus := range list {
		bus.DateJson, bus.Time = utill.Format(bus.Date)
		bus.PhotoUrl = utill.PutMediaDomain(bus.PhotoUrl)
	}

	return list, nil
}

func (s *Booking) Update(data *dtos.Booking, userId int) (*dtos.Booking, error) {
	ok, err := s.Repo.Booking.ExistsByIdAndPartyId(data.ID, userId)
	if err != nil || !ok {
		return nil, errors.New("access denied")
	}

	dataModel := data.MapToModel()
	if dataModel.Date.Before(time.Now()) {
		return nil, errors.New("booking time is past or not given")
	}

	model, err := s.Repo.Booking.GetById(data.ID)

	model.Date = dataModel.Date

	model, err = s.Repo.Booking.Update(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)
	return data, nil
}

func (s *Booking) DeleteById(id int, userId int) error {
	ok, err := s.Repo.Booking.ExistsByIdAndPartyId(id, userId)
	if err != nil || !ok {
		return errors.New("access denied")
	}

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
