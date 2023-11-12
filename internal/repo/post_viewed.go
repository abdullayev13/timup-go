package repo

import (
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
	"time"
)

type PostViewed struct {
	DB *gorm.DB
}

func (r *PostViewed) Create(model *models.PostViewed) (*models.PostViewed, error) {
	if model.CreatedAt.IsZero() {
		model.CreatedAt = time.Now()
	}

	err := r.DB.Create(model).Error

	return model, err
}
