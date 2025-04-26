package handlers

import (
	"net/http"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/internal/services"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/messages"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/response"
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
	var params models.HotelFilterParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.WithError(ctx, http.StatusBadRequest, messages.InvalidJSONOrMissingFields, err)
		return
	}
	params.NormalizeFeatures()

	ctx.JSON(200, params)
}

func (h *HotelHandler) Hotel(ctx *gin.Context) {

}

func (h *HotelHandler) Search(ctx *gin.Context) {
}
