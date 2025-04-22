package handlers

import (
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/services"
	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	hotelService *services.HotelService
}

func NewHotelHandler(hotelService *services.HotelService) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
	}
}
func (h *HotelHandler) Hotels(ctx *gin.Context) {
	// Read filter,pagination, sort query params from request

	// Get hotels from hotel service

	// Return hotels as response
	// c.JSON(http.StatusOK, hotels)
}

func (h *HotelHandler) Hotel(ctx *gin.Context) {
}

func (h *HotelHandler) Search(ctx *gin.Context) {
}
