package handlers

import (
	"net/http"

	"nganterin-go/dto"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterHotel(c *gin.Context) {
	var hotelInput dto.HotelInputDTO

	if err := c.ShouldBindJSON(&hotelInput); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid input",
			Error:   err.Error(),
		})
		return
	}

	hotelID, err := h.service.RegisterHotel(hotelInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create hotel",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "Hotel created successfully",
		Data:    hotelID,
	})
}

func (h *compHandlers) GetAllHotels(c *gin.Context) {
	result, err := h.service.GetAllHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		})
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Get all hotels successfully",
		Data:    result,
	})
}
