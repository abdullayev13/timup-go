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
		return nil, fmt.Errorf("error uploading photo: %s", err.Error())
	}

	data.MediaType = models.Photo
	if data.Video != nil {
		data.VideoPath, err = utill.Upload(data.Video, config.PostDir)
		if err != nil {
			return nil, fmt.Errorf("error uploading video: %s", err.Error())
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
	res.Photo = utill.PutMediaPostDomain(res.Photo)
	res.Video = utill.PutMediaPostDomain(res.Video)

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

		post.Photo = utill.PutMediaPostDomain(post.Photo)
		post.Video = utill.PutMediaPostDomain(post.Video)
		listDto[i] = post
	}

	return listDto, nil
}

func (s *Post) GetDetail(id, userId int) (*dtos.PostDetail, error) {
	dto, err := s.Repo.Post.GetDetail(id)
	if err != nil {
		return nil, err
	}
	{
		pv := new(models.PostViewed)
		pv.PostId = dto.Id
		pv.ViewerUserId = userId
		go s.Repo.PostViewed.Create(pv)
	}

	dto.PhotoPath = utill.PutMediaPostDomain(dto.PhotoPath)
	dto.VideoPath = utill.PutMediaPostDomain(dto.VideoPath)
	dto.PosterPhotoUrl = utill.PutMediaDomain(dto.PosterPhotoUrl)

	return dto, nil
}

func (s *Post) Update(data *dtos.PostFile, userId int) (*dtos.Post, error) {
	if !s.Repo.Business.ExistsByIdAndUserId(data.BusinessId, userId) {
		return nil, errors.New("permission denied: you are not owner of this business profile")
	}
	model, err := s.Repo.Post.GetById(data.Id)
	if err != nil {
		return nil, err
	}

	if model.BusinessId != data.BusinessId {
		return nil, errors.New("access denied")
	}

	if data.MediaType == "" {
		if data.Video == nil && model.VideoPath == "" {
			data.MediaType = models.Photo
		} else {
			data.MediaType = models.Video
		}
	}

	switch data.MediaType {
	case models.Photo:
		model.VideoPath = ""
		model.MediaType = string(models.Photo)
	case models.Video:
		if data.Video == nil || model.VideoPath == "" {
			return nil, errors.New("video not found")
		}
		model.MediaType = string(models.Video)
	default:
		return nil, errors.New("media_type not found")
	}

	if data.Photo != nil {
		model.PhotoPath, err = utill.Upload(data.Photo, config.PostDir)
		if err != nil {
			return nil, fmt.Errorf("error uploading photo: %s", err.Error())
		}
	}

	if data.Video != nil {
		data.VideoPath, err = utill.Upload(data.Video, config.PostDir)
		if err != nil {
			return nil, fmt.Errorf("error uploading video: %s", err.Error())
		}
		model.MediaType = string(models.Video)
	}

	model, err = s.Repo.Post.Update(model)
	if err != nil {
		return nil, err
	}

	res := new(dtos.Post)
	res.MapFromModel(model)
	res.Photo = utill.PutMediaPostDomain(res.Photo)
	res.Video = utill.PutMediaPostDomain(res.Video)

	return res, nil
}

func (s *Post) DeleteById(id int, userId int) error {
	model, err := s.Repo.Post.GetById(id)
	if err != nil {
		return err
	}

	if !s.Repo.Business.ExistsByIdAndUserId(model.BusinessId, userId) {
		return errors.New("access denied")
	}

	return s.Repo.Post.DeleteById(id)
}

// other

func (s *Post) GetListFollowed(data *dtos.PostFilter, userId int) ([]*dtos.PostDetail, error) {
	if data.Limit == 0 {
		data.Limit = 100
	}

	list, err := s.Repo.Post.GetListFollowed(data, userId)
	if err != nil {
		return nil, err
	}

	for _, model := range list {
		model.PhotoPath = utill.PutMediaPostDomain(model.PhotoPath)
		model.VideoPath = utill.PutMediaPostDomain(model.VideoPath)
		model.PosterPhotoUrl = utill.PutMediaPostDomain(model.PosterPhotoUrl)
	}

	return list, nil
}
