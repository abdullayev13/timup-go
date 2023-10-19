package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
	"errors"
)

type Business struct {
	Repo *repo.Repo
}

func (s *Business) Create(data *dtos.BusinessProfile) (*dtos.BusinessProfile, error) {
	category, err := s.Repo.Category.GetById(data.CategoryId)
	if err != nil {
		return nil, errors.New("category not found")
	}
	if category.ParentId == 0 {
		return nil, errors.New("category not valid")
	}

	model := data.MapToModel()
	model, err = s.Repo.Business.Create(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)
	data.SetCategoryName(category.Name)

	return data, nil
}

func (s *Business) GetByUserId(userId int) (*dtos.BusinessProfile, error) {
	model, err := s.Repo.Business.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	var categoryName string
	category, err := s.Repo.Category.GetById(model.WorkCategoryId)
	if err == nil {
		categoryName = category.Name
	}

	dto := new(dtos.BusinessProfile)
	dto.MapFromModel(model)
	dto.SetCategoryName(categoryName)

	return dto, nil
}

func (s *Business) GetByGetByCategory(data *dtos.BusinessFilter) ([]*dtos.BusinessMini, error) {
	if data.Limit == 0 {
		data.Limit = 100
	}
	if data.CategoryId == 0 {
		return nil, errors.New("category not given")
	}

	list, err := s.Repo.Business.GetByCategory(data)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].PhotoUrl = utill.PutMediaDomain(list[i].PhotoUrl)
	}

	return list, nil
}

func (s *Business) Update(dto *dtos.BusinessProfile) (*dtos.BusinessProfile, error) {
	category, err := s.Repo.Category.GetById(dto.CategoryId)
	if err != nil {
		return nil, errors.New("category not found")
	}
	if category.ParentId == 0 {
		return nil, errors.New("category not valid")
	}

	model := dto.MapToModel()
	orgModel, err := s.Repo.Business.GetById(model.ID)
	if err != nil {
		return nil, err
	}

	if orgModel.UserID != model.UserID {
		return nil, errors.New("access denied")
	}

	model, err = s.Repo.Business.Update(model)
	if err != nil {
		return nil, err
	}

	dto.MapFromModel(model)
	dto.SetCategoryName(category.Name)

	return dto, nil
}

func (s *Business) DeleteByUserId(userId int) error {
	return s.Repo.Business.DeleteByUserId(userId)
}

//	other

func (s *Business) GetProfileById(id int) (*dtos.BusinessData, error) {
	model, err := s.Repo.Business.GetById(id)
	if err != nil {
		return nil, err
	}

	var categoryName string
	category, err := s.Repo.Category.GetById(model.WorkCategoryId)
	if err == nil {
		categoryName = category.Name
	}

	user, err := s.Repo.Users.GetById(model.UserID)
	if err != nil {
		return nil, err
	}

	dto := new(dtos.BusinessData)
	dto.MapFromModel(model)
	dto.SetCategoryName(categoryName)
	dto.SetUser(user)

	return dto, nil
}
