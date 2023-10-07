package dtos

import "mime/multipart"

type Sign struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type PhotoReq struct {
	ProfilePhoto *multipart.FileHeader `form:"profile_photo" json:"profile_photo"`
	PhotoUrl     string                `form:"profile_url" json:"profile_url"`
	UserId       int                   `form:"-" json:"-"`
}

type PhotoRes struct {
	PhotoUrl string `form:"profile_url" json:"profile_url"`
}
