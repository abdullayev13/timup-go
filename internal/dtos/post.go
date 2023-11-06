package dtos

import (
	"abdullayev13/timeup/internal/models"
	"mime/multipart"
	"time"
)

type Post struct {
	Id          int              `json:"id" form:"id"`
	MediaType   models.MediaType `json:"media_type" form:"media_type"`
	Photo       string           `json:"photo" form:"photo"`
	Video       string           `json:"video,omitempty" form:"video,omitempty"`
	Title       string           `json:"title" form:"title"`
	Description string           `json:"description" form:"description"`
	BusinessId  int              `json:"business_id,omitempty" form:"business_id,omitempty"`
	CreatedAt   time.Time        `json:"created_at,omitempty" form:"created_at,omitempty"`
}

type PostFile struct {
	Id          int                   `form:"id"`
	MediaType   models.MediaType      `form:"media_type"`
	Photo       *multipart.FileHeader `form:"photo"`
	Video       *multipart.FileHeader `form:"video"`
	PhotoPath   string                `form:"-"`
	VideoPath   string                `form:"-"`
	Title       string                `form:"title"`
	Description string                `form:"description"`
	BusinessId  int                   `form:"business_id"`
	CreatedAt   time.Time             `form:"created_at"`
}

type PostFilter struct {
	Limit      int
	Offset     int
	BusinessId int
	MediaType  models.MediaType
}

type PostDetail struct {
	Id             int              `form:"id"`
	MediaType      models.MediaType `form:"media_type"`
	Photo          string           `form:"photo"`
	Video          string           `form:"video"`
	Title          string           `form:"title"`
	Description    string           `form:"description"`
	BusinessId     int              `form:"business_id"`
	CreatedAt      time.Time        `form:"created_at"`
	PosterPhotoUrl string           `form:"poster_photo_url"`
	PosterName     string           `form:"poster_name"`
}

func (d *PostFile) MapToModel() *models.Post {
	m := new(models.Post)

	m.Id = d.Id
	m.PhotoPath = d.PhotoPath
	m.VideoPath = d.VideoPath
	m.Title = d.Title
	m.Description = d.Description
	m.BusinessId = d.BusinessId
	m.CreatedAt = d.CreatedAt
	m.MediaType = d.MediaType

	return m
}

func (d *Post) MapFromModel(m *models.Post) *Post {

	d.Id = m.Id
	d.Photo = m.PhotoPath
	d.Video = m.VideoPath
	d.Title = m.Title
	d.Description = m.Description
	d.BusinessId = m.BusinessId
	d.CreatedAt = m.CreatedAt
	d.MediaType = m.MediaType

	return d
}
