package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/repo"
	"errors"
)

type BookingCategory struct {
	Repo *repo.Repo
}

func (s *BookingCategory) Create(data *dtos.BookingCategory, userId int) (*dtos.BookingCategory, error) {
	if !s.Repo.Business.ExistsByIdAndUserId(data.BusinessId, userId) {
		return nil, errors.New("business_id incorrect or not found")
	}

	data.Id = 0
	model := data.MapToModel()

	model, err := s.Repo.BookingCategory.Create(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)

	return data, nil
}

func (s *BookingCategory) GetByBusinessId(businessId int) ([]*dtos.BookingCategory, error) {
	modelSlc, err := s.Repo.BookingCategory.GetByBusinessId(businessId)
	if err != nil {
		return nil, err
	}

	dtoSlc := make([]*dtos.BookingCategory, len(modelSlc))

	for i, category := range modelSlc {
		dto := new(dtos.BookingCategory)
		dto.MapFromModel(category)

		dtoSlc[i] = dto
	}

	return dtoSlc, nil
}

func (s *BookingCategory) DeleteById(id, userId int) error {
	model, err := s.Repo.BookingCategory.GetById(id)
	if err != nil {
		return err
	}

	if !s.Repo.Business.ExistsByIdAndUserId(model.BusinessId, userId) {
		return errors.New("access denied")
	}

	return s.Repo.BookingCategory.DeleteById(id)
}
