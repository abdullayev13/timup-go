package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/repo"
)

type Business struct {
	Repo *repo.Repo
}

func (s *Business) Create(data *dtos.BusinessProfile) (*dtos.BusinessProfile, error) {
	model := data.MapToModel()

	model, err := s.Repo.Business.Create(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)

	return data, nil
}

func (s *Business) GetByUserId(userId int) (*dtos.BusinessProfile, error) {
	model, err := s.Repo.Business.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	dto := new(dtos.BusinessProfile)
	dto.MapFromModel(model)

	return dto, nil
}

func (s *Business) Update(dto *dtos.BusinessProfile) (*dtos.BusinessProfile, error) {
	model := dto.MapToModel()
	_, err := s.Repo.Business.GetById(model.ID)
	if err != nil {
		return nil, err
	}

	model, err = s.Repo.Business.Update(model)
	if err != nil {
		return nil, err
	}

	dto.MapFromModel(model)

	return dto, nil
}

func (s *Business) DeleteByUserId(userId int) error {
	return s.Repo.Business.DeleteByUserId(userId)
}
