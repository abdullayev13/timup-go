package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/repo"
	"errors"
)

type Category struct {
	Repo *repo.Repo
}

func (s *Category) Create(data *dtos.WorkCategory) (*dtos.WorkCategory, error) {
	data.ID = 0
	model := data.MapToModel()

	model, err := s.Repo.Category.Create(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)

	return data, nil
}

func (s *Category) GetByParentId(userId int) ([]*dtos.WorkCategory, error) {
	modelSlc, err := s.Repo.Category.GetByParentId(userId)
	if err != nil {
		return nil, err
	}

	dtoSlc := make([]*dtos.WorkCategory, len(modelSlc))

	for i, category := range modelSlc {
		dto := new(dtos.WorkCategory)
		dto.MapFromModel(category)

		dtoSlc[i] = dto
	}

	return dtoSlc, nil
}

func (s *Category) Update(data *dtos.WorkCategory) (*dtos.WorkCategory, error) {
	model, err := s.Repo.Category.GetById(data.ID)
	if err != nil {
		return nil, err
	}

	model.Name = data.Name
	if data.ParentId != 0 && model.ParentId != data.ParentId {
		_, err = s.Repo.Category.GetById(data.ParentId)
		if err != nil {
			return nil, errors.New("error with parent category: " + err.Error())
		}

		model.ParentId = data.ParentId
	}

	model, err = s.Repo.Category.Update(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)
	return data, nil
}

func (s *Category) DeleteById(id int) error {
	return s.Repo.Category.DeleteById(id)
}
