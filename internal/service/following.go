package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
	"errors"
	"time"
)

type Following struct {
	Repo *repo.Repo
}

func (s *Following) Create(data *dtos.Following) (*dtos.Following, error) {
	sameOwner := s.Repo.Business.ExistsByIdAndUserId(data.BusinessId, data.FollowerId)
	if sameOwner {
		return nil, errors.New("following your self not allowed")
	}

	data.CreatedAt = time.Now()

	model := data.MapToModel()

	var err error
	model, err = s.Repo.Following.Create(model)
	if err != nil {
		return nil, err
	}

	data.MapFromModel(model)

	return data, nil
}

func (s *Following) DeleteById(id int) error {
	return s.Repo.Following.DeleteById(id)
}

func (s *Following) DeleteByFollower(businessId, followerId int) error {
	err := s.Repo.Following.Delete(businessId, followerId)
	return err
}

func (s *Following) GetBusinessList(data *dtos.FollowedFilter) ([]*dtos.BusinessLI, error) {
	if data.Limit == 0 {
		data.Limit = 100
	}
	list, err := s.Repo.Following.GetBusinessList(data)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].PhotoUrl = utill.PutMediaDomain(list[i].PhotoUrl)
	}

	return list, nil
}
