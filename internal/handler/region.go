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
	"Andijon viloyati",
	"Andijon shahri",
	"Buxoro viloyati",
	"Buxoro shahri",
	"Fargʻona viloyati",
	"Fargʻona shahri",
	"Jizzax viloyati",
	"Jizzax shahri",
	"Xorazm viloyati",
	"Urganch shahri",
	"Namangan viloyati",
	"Namangan shahri",
	"Navoiy viloyati",
	"Navoiy shahri",
	"Qashqadaryo viloyati",
	"Qarshi shahri",
	"Qoraqalpogʻiston Respublikasi",
	"Nukus shahri",
	"Samarqand viloyati",
	"Samarqand shahri",
	"Sirdaryo viloyati",
	"Guliston shahri",
	"Surxondaryo viloyati",
	"Termiz shahri",
	"Toshkent viloyati",
	"Toshkent shahri",
}
