package controllers

import (
	"net/http"
	"nganterin-go/api/reservations/services"
	"nganterin-go/api/reviews/dto"
	"nganterin-go/pkg/exceptions"
	"nganterin-go/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) FindByUserID(ctx *gin.Context) {
	userData, err := helpers.GetUserData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	result, err := h.services.FindByUserID(ctx, userData.ID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Data:    result,
		Message: "data retrieved successfully",
	})
}

func (h *CompControllersImpl) FindByHotelID(ctx *gin.Context) {
	hotelID := ctx.Query("id")

	if hotelID == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.FindByHotelID(ctx, hotelID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Data:    result,
		Message: "data retrieved successfully",
	})
}

func (h *CompControllersImpl) FindByReservationKey(ctx *gin.Context) {
	reservationKey := ctx.Query("key")

	if reservationKey == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	result, err := h.services.FindByReservationKey(ctx, reservationKey)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Data:    result,
		Message: "data retrieved successfully",
	})
}

func (h *CompControllersImpl) CheckIn(ctx *gin.Context) {
	reservationKey := ctx.Query("key")

	if reservationKey == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.CheckIn(ctx, reservationKey)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "reservation checked in successfully",
	})
}

func (h *CompControllersImpl) CheckOut(ctx *gin.Context) {
	reservationKey := ctx.Query("key")

	if reservationKey == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.CheckOut(ctx, reservationKey)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "reservation checked out successfully",
	})
}

func (h *CompControllersImpl) YearlyReservationAnalytic(ctx *gin.Context) {
	partnerData, err := helpers.GetPartnerData(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	result, err := h.services.YearlyReservationAnalytic(ctx, partnerData.ID)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "data retrieved successfully",
		Data:    result,
	})
}
