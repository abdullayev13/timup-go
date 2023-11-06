package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/repo"
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

func (s *Category) Update(dto *dtos.WorkCategory) (*dtos.WorkCategory, error) {
	return nil, nil
}

func (s *Category) DeleteById(id int) error {
	return s.Repo.Category.DeleteById(id)
}
