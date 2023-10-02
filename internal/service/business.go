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
	//model := dto.MapToUser()
	//orgModel, err := s.Repo.Users.GetById(model.ID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//model.PhoneNumber = orgModel.PhoneNumber
	//if model.PhotoUrl == "" {
	//	model.PhotoUrl = orgModel.PhotoUrl
	//}
	//
	//model, err = s.Repo.Users.Update(model)
	//if err != nil {
	//	return nil, err
	//}
	//
	//dto.MapFromUser(model)
	//
	//return dto, nil

	return nil, nil
}

func (s *Business) DeleteByUserId(userId int) error {
	return s.Repo.Business.DeleteByUserId(userId)
}
