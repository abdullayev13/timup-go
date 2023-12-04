package service

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"abdullayev13/timeup/internal/pkg/transcode_upload"
	"abdullayev13/timeup/internal/repo"
	"abdullayev13/timeup/internal/utill"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Post struct {
	Repo *repo.Repo
}

func (s *Post) Create(data *dtos.PostFile, userId int) (map[string]any, error) {
	if !s.Repo.Business.ExistsByIdAndUserId(data.BusinessId, userId) {
		return nil, errors.New("permission denied: you are not owner of this business profile")
	}

	var (
		videoErr, imgErr error
		spendingMinutes  int64
	)
	var wg sync.WaitGroup
	if data.Photo != nil {
		spendingMinutes += (data.Photo.Size >> 20) / 5

		wg.Add(1)
		go transcode_upload.TranscodeAndUploadS3Img(data.Photo, func(url string, err error) {
			defer wg.Done()
			if err != nil {
				imgErr = err
				return
			}
			data.PhotoPath = url
		})
	}

	data.MediaType = models.Photo
	if data.Video != nil {
		spendingMinutes += (data.Video.Size >> 20) / 7

		wg.Add(1)
		go transcode_upload.TranscodeAndUploadS3Video(data.Video, func(url string, err error) {
			defer wg.Done()
			if err != nil {
				videoErr = err
				return
			}
			data.VideoPath = url
		})

		data.MediaType = models.Video
	}

	data.Id = 0
	data.CreatedAt = time.Now()

	go func() {
		wg.Wait()

		if imgErr != nil {
			fmt.Printf("error uploading img[user_id=%d]:%s", userId, imgErr.Error())
			return
		}
		if videoErr != nil {
			fmt.Printf("error uploading video[user_id=%d]:%s", userId, videoErr.Error())
			return
		}

		model := data.MapToModel()

		model, err := s.Repo.Post.Create(model)
		if err != nil {
			jsonData, _ := json.Marshal(model)
			fmt.Printf("error creating post[user_id=%d]:%s; %x", userId, err.Error(), jsonData)
			return
		}
	}()

	return map[string]any{
		"spending_minute": spendingMinutes,
	}, nil

	//var err error
	//if data.Photo != nil {
	//	data.PhotoPath, err = utill.TranscodeAndUploadS3Img(data.Photo)
	//	if err != nil {
	//		return nil, fmt.Errorf("error uploading photo: %s", err.Error())
	//	}
	//}
	//
	//data.MediaType = models.Photo
	//if data.Video != nil {
	//	data.VideoPath, err = utill.TranscodeAndUploadS3Video(data.Video)
	//	if err != nil {
	//		return nil, fmt.Errorf("error uploading video: %s", err.Error())
	//	}
	//	data.MediaType = models.Video
	//}
	//
	//data.Id = 0
	//data.CreatedAt = time.Now()
	//
	//model := data.MapToModel()
	//
	//model, err = s.Repo.Post.Create(model)
	//if err != nil {
	//	return nil, err
	//}
	//
	//res := new(dtos.Post)
	//res.MapFromModel(model)
	//res.Photo = utill.PutMediaPostDomain(res.Photo)
	//res.Video = utill.PutMediaPostDomain(res.Video)
	//
	//return res, nil
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
			model.MediaType = string(models.Photo)
		} else {
			data.MediaType = models.Video
			model.MediaType = string(models.Video)
		}
	}
	if data.Title != "" {
		model.Title = data.Title
	}
	if data.Description != "" {
		model.Description = data.Description
	}

	switch data.MediaType {
	case models.Photo:
		model.VideoPath = ""
		model.MediaType = string(models.Photo)
	case models.Video:
		if data.Video == nil && model.VideoPath == "" {
			return nil, errors.New("video not found")
		}
		model.MediaType = string(models.Video)
	default:
		return nil, errors.New("media_type not found")
	}

	if data.Photo != nil {
		model.PhotoPath, err = utill.TranscodeAndUploadS3Img(data.Photo)
		if err != nil {
			return nil, fmt.Errorf("error uploading photo: %s", err.Error())
		}
	}

	if data.Video != nil {
		data.VideoPath, err = utill.TranscodeAndUploadS3Video(data.Video)
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

	for _, dto := range list {
		dto.PhotoPath = utill.PutMediaPostDomain(dto.PhotoPath)
		dto.VideoPath = utill.PutMediaPostDomain(dto.VideoPath)
		dto.PosterPhotoUrl = utill.PutMediaPostDomain(dto.PosterPhotoUrl)
	}

	return list, nil
}
