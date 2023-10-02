package handler

import (
	"abdullayev13/timeup/internal/handler/response"
	"github.com/gin-gonic/gin"
)

type Region struct {
}

func (h *Region) Get(c *gin.Context) {

	response.Success(c, region)
}

var region = []string{
	"Tashkent viloyat", "Tashkent",
	"Andijan viloyat", "Andijan",
	"Bukhara viloyat", "Bukhara",
	"Fergana viloyat", "Fergana",
	"Jizzakh viloyat", "Jizzakh",
	"Namangan viloyat", "Namangan",
	"Navoiy viloyat", "Navoiy",
	"Qashqadaryo viloyat", "Qarshi",
	"Samarqand viloyat", "Samarkand",
	"Sirdaryo viloyat", "Guliston",
	"Surxondaryo viloyat", "Termez",
	"Tashkent viloyat", "Nurafshon",
	"Xorazm viloyat", "Urgench",
	"Republic of Karakalpakstan", "Nukus",
}
