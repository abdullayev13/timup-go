package repo

import (
	"abdullayev13/timeup/internal/models"
	"gorm.io/gorm"
)

type Users struct {
	DB *gorm.DB
}

//func (s *Users) SignUp(sign endpoint.Sign) (*moduls.User, error) {
//	exists := s.existsBYUserName(sign.UserName)
//	if exists {
//		return nil, errors.New("username exists")
//	}
//	user := moduls.User{UserName: sign.UserName, Password: sign.Password, Name: sign.Name}
//	s.DB.Create(&user)
//	return &user, nil
//}
//
//func (s *Users) LogIn(sign endpoint.Sign) (string, error) {
//	exists := s.existsBYUserName(sign.UserName)
//	if !exists {
//		return "", errors.New("username or password wrong")
//	}
//	var userId int
//	s.DB.Model(&moduls.User{}).
//		Select("id").
//		Where("user_name = ?", sign.UserName).
//		Find(&userId)
//	token, err := s.TokenJWT.GenerateToken(userId)
//	if err != nil {
//		return "", err
//	}
//	return token, nil
//}
//
//func (s *Users) AddPost(post *moduls.Post) error {
//	err := s.DB.Create(post).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (s *Users) PostsOfUser(userId int) []moduls.Post {
//	var posts []moduls.Post
//	s.DB.Model(&moduls.Post{}).
//		Where("user_id = ?", userId).
//		Find(&posts)
//	return posts
//}
//func (s *Users) Like(like *moduls.Like) error {
//	var exists bool
//	s.DB.Model(&moduls.Like{}).
//		Select("count(*) > 0").
//		Where("user_id = ? AND liked_id = ? AND type = ?", like.UserID, like.LikedID, like.Type).
//		Find(&exists)
//	if exists {
//		return errors.New("already liked")
//	}
//	{ // chack type and likeId are correct
//		var module interface{}
//		switch like.Type {
//		case "post":
//			module = &moduls.Post{}
//		case "comment":
//			module = &moduls.Comment{}
//		default:
//			return errors.New("type should be post or comment")
//		}
//		exists := s.exists(module, " id = ?", like.LikedID)
//		if !exists {
//			return errors.New(like.Type + " not exists by likedId")
//		}
//	}
//	s.DB.Create(like)
//	return nil
//}
//
//func (s *Users) UserIdFromToken(tokenStr string) (int, error) {
//	return s.TokenJWT.ParseToken(tokenStr)
//}
//
//func (s *Users) AddComment(comment *moduls.Comment) error {
//	exists := s.existsById(moduls.Post{}, comment.PostID)
//	if !exists {
//		return errors.New("post does not exist by id")
//	}
//	s.DB.Create(comment)
//	return nil
//}
//
//func (s *Users) CommentsByPostId(postId int) ([]moduls.Comment, error) {
//	exists := s.existsById(moduls.Post{}, postId)
//	if !exists {
//		return nil, errors.New("post does not exist by id")
//	}
//	var comments []moduls.Comment
//	s.DB.Where("post_id = ?", postId).
//		Find(&comments)
//	return comments, nil
//}

func (r *Users) Create(model *models.User) (*models.User, error) {
	err := r.DB.Create(model).Error

	return model, err
}

func (r *Users) GetByPhoneNumber(phoneNumber string) (*models.User, error) {
	model := new(models.User)
	err := r.DB.Where("phone_number = ?", phoneNumber).
		First(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Users) GetById(id int) (*models.User, error) {
	model := new(models.User)
	err := r.DB.First(model, id).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Users) Update(model *models.User) (*models.User, error) {
	err := r.DB.Save(model).Error

	return model, err
}

func (r *Users) DeleteById(id int) error {
	return r.DB.Delete(models.User{ID: id}, id).Error
}

// special funcs
func (r *Users) ExistsByUsername(username string) bool {
	exists := r.exists(&models.User{}, "username=?", username)
	//r.DB.Model(&models.User{}).
	//	Select("count(*) > 0").
	//	Where("user_name = ?", username).
	//	Find(&exists)
	return exists
}

func (r *Users) ExistsByPhoneNumber(phoneNumber string) bool {
	exists := r.exists(&models.User{}, "phone_number=?", phoneNumber)
	return exists
}

func (r *Users) exists(module any, whereQuery string, args ...any) bool {
	var exists bool
	r.DB.Model(module).
		Select("count(*) > 0").
		Where(whereQuery, args...).
		Find(&exists)
	return exists
}

func (r *Users) ExistsById(module any, id any) bool {
	return r.exists(module, "id = ?", id)
}
