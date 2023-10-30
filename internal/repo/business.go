package repo

import (
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/models"
	"errors"
	"fmt"
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

func (r *Business) GetProfileById(id int) (*dtos.BusinessData, error) {
	fullData, err := r.GetFullDataBy("b.id=?", id)
	if err != nil {
		return nil, err
	}

	if id != fullData.ID {
		return nil, errors.New("not found")
	}

	dto := new(dtos.BusinessData)
	{
		dto.ID = fullData.ID
		dto.UserID = fullData.UserID
		dto.CategoryId = fullData.CategoryId
		dto.CategoryName = fullData.CategoryName
		dto.OfficeAddress = fullData.OfficeAddress
		dto.OfficeName = fullData.OfficeName
		dto.Experience = fullData.Experience
		dto.Bio = fullData.Bio
		dto.DayOffs = fullData.DayOffs
		dto.FistName = fullData.FistName
		dto.LastName = fullData.LastName
		dto.UserName = fullData.UserName
		dto.PhoneNumber = fullData.PhoneNumber
		dto.Address = fullData.Address
		dto.PhotoUrl = fullData.PhotoUrl
		dto.FollowersCount = fullData.FollowersCount
	}

	return dto, nil
}

func (r *Business) GetProfileByUserId(userId int) (*dtos.BusinessFullData, error) {
	fullData, err := r.GetFullDataBy("u.id=?", userId)
	if err != nil {
		return nil, err
	}

	return fullData, nil
}

func (r *Business) GetFullDataBy(whereQuery string, args ...any) (*dtos.BusinessFullData, error) {
	dto := new(dtos.BusinessFullData)

	err := r.DB.Raw(
		fmt.Sprintf(`SELECT b.id,
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
       c.id      as category_id,
       c.name      as category_name,
(SELECT count(f.id) FROM followings f WHERE f.business_id = b.id) as followers_count
FROM business_profiles b
         JOIN users u on b.user_id = u.id
         JOIN work_categories c on b.work_category_id = c.id
WHERE %s
GROUP BY b.id, c.id, u.id`, whereQuery), args...).
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
