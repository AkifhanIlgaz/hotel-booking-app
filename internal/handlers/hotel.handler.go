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
	hotels, err := h.hotelService.GetHotels()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, hotels)
}

func (h *HotelHandler) Hotel(ctx *gin.Context) {

}

func (h *HotelHandler) Search(ctx *gin.Context) {
}
