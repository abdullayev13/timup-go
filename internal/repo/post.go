package repo

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
	"strings"
)

type Post struct {
	DB *gorm.DB
}

func (r *Post) Create(model *models.Post) (*models.Post, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *Post) GetById(id int) (*models.Post, error) {
	model := new(models.Post)
	err := r.DB.First(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Post) GetList(data *dtos.PostFilter) ([]*models.Post, error) {
	tx := r.DB.Order("date")

	if data.Offset != 0 {
		tx.Offset(data.Offset)
	}
	if data.Limit != 0 {
		tx.Limit(data.Limit)
	}
	if data.BusinessId != 0 {
		tx.Where("business_id = ?", data.BusinessId)
	}
	if data.MediaType != "" {
		tx.Where("media_type = ?", data.MediaType)
	}

	models := make([]*models.Post, 0)

	err := tx.Find(&models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (r *Post) Update(model *models.Post) (*models.Post, error) {
	err := r.DB.Save(model).Error

	return model, err
}

func (r *Post) DeleteById(id int) error {
	model := new(models.Post)
	err := r.DB.Where("id = ?", id).Delete(model).Error

	return err
}

// other

func (r *Post) GetDetail(id int) (*dtos.PostDetail, error) {
	model := new(dtos.PostDetail)
	err := r.DB.Raw(`SELECT p.*,
       u.photo_url                as poster_photo_url,
       u.fist_name ||' '|| u.last_name as poster_name
FROM posts p
         JOIN business_profiles b on p.business_id = b.id
         JOIN users u on b.user_id = u.id
WHERE p.id = ?`, id).Find(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Post) GetMiniList(data *dtos.PostFilter) ([]*models.Post, error) {
	query := []string{`SELECT id,
       media_type,
       photo_path,
       video_path,
       title,
       CASE
           WHEN LENGTH(description) < 64 THEN description
           ELSE SUBSTRING(description, 1, 61) || '...'
           END AS description,
       business_id
FROM posts p WHERE TRUE`}
	args := []any{}

	if data.BusinessId != 0 {
		query = append(query, "AND business_id = ?")
		args = append(args, data.BusinessId)
	}
	if data.MediaType != "" {
		query = append(query, "AND media_type > ?")
		args = append(args, data.MediaType)
	}
	query = append(query, "ORDER BY created_at LIMIT ? OFFSET ?")
	args = append(args, data.Limit, data.Offset)

	res := make([]*models.Post, 0, data.Limit)
	err := r.DB.Raw(strings.Join(query, " "), args...).Find(&res).Error

	return res, err
}

func (r *Post) GetListFollowed(data *dtos.PostFilter, userId int) ([]*dtos.PostDetail, error) {
	query := []string{`SELECT p.*,
       u.photo_url                       as poster_photo_url,
       u.fist_name || ' ' || u.last_name as poster_name
FROM posts p
         JOIN business_profiles b on p.business_id = b.id
         JOIN users u on b.user_id = u.id
         JOIN followings f on f.business_id = b.id
WHERE f.follower_id = ?`}
	args := []any{userId}

	if data.MediaType != "" {
		query = append(query, "AND p.media_type > ?")
		args = append(args, data.MediaType)
	}
	query = append(query, "ORDER BY p.created_at DESC LIMIT ? OFFSET ?")
	args = append(args, data.Limit, data.Offset)

	res := make([]*dtos.PostDetail, 0, data.Limit)
	err := r.DB.Raw(strings.Join(query, " "), args...).Find(&res).Error

	return res, err
}
