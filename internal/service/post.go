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

type Post struct {
	Repo *repo.Repo
}

func (s *Post) Create(data *dtos.PostFile, userId int) (*dtos.Post, error) {
	if !s.Repo.Business.ExistsByIdAndUserId(data.BusinessId, userId) {
		return nil, errors.New("permission denied: you are not owner of this business profile")
	}

	if data.Photo == nil {
		return nil, errors.New("photo not found")
	}

	var err error
	data.PhotoPath, err = utill.Upload(data.Photo, config.PostDir)
	if err != nil {
		return nil, fmt.Errorf("error uploading photo: ", err.Error())
	}

	data.MediaType = models.Photo
	if data.Video != nil {
		data.VideoPath, err = utill.Upload(data.Video, config.PostDir)
		if err != nil {
			return nil, fmt.Errorf("error uploading video: ", err.Error())
		}
		data.MediaType = models.Video
	}

	data.Id = 0
	data.CreatedAt = time.Now()

	model := data.MapToModel()

	model, err = s.Repo.Post.Create(model)
	if err != nil {
		return nil, err
	}

	res := new(dtos.Post)
	res.MapFromModel(model)
	res.Photo = utill.PutMediaDomain(res.Photo)

	return res, nil
}

func (s *Post) GetList(data *dtos.PostFilter) ([]*dtos.Post, error) {
	if data.Limit == 0 {
		data.Limit = 100
	}

	list, err := s.Repo.Post.GetMiniList(data)
	if err != nil {
		return nil, err
	}

	listDto := make([]*dtos.Post, len(list))
	for i, model := range list {
		post := new(dtos.Post).
			MapFromModel(model)

		post.Photo = utill.PutMediaDomain(post.Photo)
		listDto[i] = post
	}

	return listDto, nil
}

func (s *Post) GetById(id int) (*dtos.PostDetail, error) {
	return s.Repo.Post.GetDetail(id)
}

func (s *Post) DeleteById(id int) error {
	return s.Repo.Post.DeleteById(id)
}
