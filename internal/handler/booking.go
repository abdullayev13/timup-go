package handler

import (
	"abdullayev13/timeup/internal/config"
	"abdullayev13/timeup/internal/dtos"
	"abdullayev13/timeup/internal/handler/response"
	"abdullayev13/timeup/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Booking struct {
	Service *service.Service
}

func (h *Booking) Create(c *gin.Context) {
	data := new(dtos.Booking)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.ClientId = c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Booking.Create(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Booking) GetList(c *gin.Context) {
	data := new(dtos.BookingFilter)
	err := c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	res, err := h.Service.Booking.GetList(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Booking) GetListByClient(c *gin.Context) {
	data := new(dtos.BookingFilter)
	err := c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)
	res, err := h.Service.Booking.GetListByClient(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Booking) GetListByBusinessId(c *gin.Context) {
	business_id, err := strconv.Atoi(c.Param("business_id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data := new(dtos.BookingFilter)
	err = c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.BusinessId = business_id
	res, err := h.Service.Booking.GetListByBusiness(data, business_id)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Booking) Update(c *gin.Context) {
	data := new(dtos.Booking)
	err := c.Bind(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	res, err := h.Service.Booking.Update(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, res)
}

func (h *Booking) DeleteByClient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	err = h.Service.Booking.DeleteByClient(id, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}

func (h *Booking) DeleteByBusiness(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	err = h.Service.Booking.DeleteByBusiness(id, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}

func (h *Booking) DeleteById(c *gin.Context) {
	userId := c.GetInt(config.UserIdKeyFromAuthMw)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	err = h.Service.Booking.DeleteById(id, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, "ok")
}

// other

func (h *Booking) GetListWithBookingCategory(c *gin.Context) {
	businessId, err := strconv.Atoi(c.Param("business_id"))
	if err != nil {
		response.FailErr(c, err)
		return
	}

	data := new(dtos.BookingFilter)
	err = c.BindQuery(data)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	userId := c.GetInt(config.UserIdKeyFromAuthMw)
	data.BusinessId = businessId

	bookings, err := h.Service.Booking.GetListByClient(data, userId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	bcs, err := h.Service.BookingCategory.GetByBusinessId(data.BusinessId)
	if err != nil {
		response.FailErr(c, err)
		return
	}

	response.Success(c, map[string]any{
		"bookings":           bookings,
		"booking_categories": bcs,
	})
}
