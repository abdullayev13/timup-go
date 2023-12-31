package repo

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"errors"
	"gorm.io/gorm"
)

type Business struct {
	DB *gorm.DB
}

func (r *Business) Create(model *models.BusinessProfile) (*models.BusinessProfile, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *Business) GetByUserId(userId int) (*models.BusinessProfile, error) {
	model := new(models.BusinessProfile)
	err := r.DB.Where("user_id = ?", userId).
		First(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Business) GetById(id int) (*models.BusinessProfile, error) {
	model := new(models.BusinessProfile)
	err := r.DB.First(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Business) GetByCategory(data *dtos.BusinessFilter) ([]*dtos.BusinessMini, error) {
	listModel := make([]*dtos.BusinessMini, 0)

	err := r.DB.Raw(`SELECT b.id as business_id,
b.user_id,
b.experience,
u.fist_name,
u.last_name,
(select f.follower_id from followings f where f.business_id = b.id
   AND f.follower_id = ?) is not null as followed,
u.photo_url
FROM users u
INNER JOIN business_profiles b on b.user_id = u.id
WHERE b.work_category_id = ? LIMIT ? OFFSET ?`, data.UserId, data.CategoryId, data.Limit, data.Offset).Find(&listModel).Error

	if err != nil {
		return nil, err
	}

	return listModel, nil
}

func (r *Business) Update(model *models.BusinessProfile) (*models.BusinessProfile, error) {
	err := r.DB.Save(model).Error

	return model, err
}

func (r *Business) DeleteById(id int) error {
	model := new(models.BusinessProfile)
	err := r.DB.Where("id = ?", id).Delete(model).Error

	return err
}

func (r *Business) DeleteByUserId(userId int) error {
	model := new(models.BusinessProfile)
	err := r.DB.Where("user_id = ?", userId).Delete(model).Error

	return err
}

//	others

func (r *Business) GetProfileById(id, viewerId int) (*dtos.BusinessData, error) {
	dto := new(dtos.BusinessData)

	err := r.DB.Raw(`SELECT b.id,
       b.office_address,
       b.office_name,
       b.experience,
       b.bio,
       b.day_offs,
       b.user_id,
       u.fist_name,
       u.user_name,
       u.last_name,
       u.phone_number,
       u.photo_url,
       u.address,
       c.id                                                              as category_id,
       c.name                                                            as category_name,
       (SELECT count(f.id) FROM followings f WHERE f.follower_id = u.id) as following_count,
       (SELECT count(f.id) FROM followings f WHERE f.business_id = b.id) as followers_count,
       exists(SELECT f.id FROM followings f WHERE f.business_id = b.id AND f.follower_id = ?) as followed,
        (SELECT count(p.id) FROM posts p WHERE p.business_id = b.id)      as posts_count
FROM business_profiles b
         JOIN users u on b.user_id = u.id
         JOIN work_categories c on b.work_category_id = c.id
WHERE b.id = ?
GROUP BY b.id, c.id, u.id`, viewerId, id).
		Find(dto).Error

	if err != nil {
		return nil, err
	}
	if dto.UserID == 0 {
		return nil, errors.New("not fount")
	}
	if id != dto.ID {
		return nil, errors.New("not found")
	}

	return dto, nil
}

func (r *Business) GetProfileByUserId(userId int) (*dtos.BusinessFullData, error) {
	dto := new(dtos.BusinessFullData)

	err := r.DB.Raw(`SELECT b.id,
       b.office_address,
       b.office_name,
       b.experience,
       b.bio,
       b.day_offs,
       u.id                                                              as user_id,
       u.fist_name,
       u.user_name,
       u.last_name,
       u.phone_number,
       u.photo_url,
       u.address,
       c.id                                                              as category_id,
       c.name                                                            as category_name,
       (SELECT COUNT(f.id) FROM followings f WHERE f.follower_id = u.id) as following_count,
       (SELECT COUNT(f.id) FROM followings f WHERE f.business_id = b.id) as followers_count,
       (SELECT count(p.id) FROM posts p WHERE p.business_id = b.id)      as posts_count
FROM users u
         LEFT JOIN business_profiles b ON b.user_id = u.id
         LEFT JOIN work_categories c ON b.work_category_id = c.id
WHERE u.id = ?
GROUP BY b.id, c.id, u.id`, userId).
		Find(dto).Error

	if err != nil {
		return nil, err
	}
	if dto.UserID == 0 {
		return nil, errors.New("not fount")
	}

	return dto, nil
}

func (r *Business) ExistsById(id int) bool {
	return r.exists("id=?", id)
}

func (r *Business) ExistsByIdAndUserId(id, userId int) bool {
	return r.exists("id=? AND user_id=?", id, userId)
}

func (r *Business) exists(whereQuery string, args ...any) bool {
	var exists bool
	r.DB.Model(&models.BusinessProfile{}).
		Select("count(*) > 0").
		Where(whereQuery, args...).
		Find(&exists)
	return exists
}
