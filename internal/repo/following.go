package repo

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
)

type Following struct {
	DB *gorm.DB
}

func (r *Following) Create(model *models.Following) (*models.Following, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *Following) GetBusinessList(data *dtos.FollowedFilter) ([]*dtos.BusinessLI, error) {
	listModel := make([]*dtos.BusinessLI, 0)

	err := r.DB.Raw(`SELECT b.id,
       c.name as category_name, 
       b.office_address, b.office_name, b.bio, b.day_offs, b.user_id,
       u.fist_name, u.user_name, u.last_name, u.phone_number, u.photo_url
FROM users u
         JOIN business_profiles b on b.user_id = u.id
         JOIN work_categories c on b.work_category_id = c.id
         JOIN followings f on f.business_id = b.id
WHERE f.follower_id = ?
LIMIT ? OFFSET ?`, data.FollowerId, data.Limit, data.Offset).
		Find(&listModel).Error

	if err != nil {
		return nil, err
	}

	return listModel, nil
}

func (r *Following) Delete(businessId, followerId int) error {
	model := new(models.Following)
	err := r.DB.Where(
		"business_id = ? And follower_id = ?",
		businessId, followerId).
		Delete(model).Error

	return err
}

func (r *Following) DeleteById(id int) error {
	model := new(models.Following)
	err := r.DB.Where("id = ?", id).Delete(model).Error

	return err
}
