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

	params.Validate()
	params.NormalizeFeatures()

	hotels, err := h.hotelService.GetHotels(params)
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, "", gin.H{
		"hotels": hotels,
	})
}

func (h *HotelHandler) Hotel(ctx *gin.Context) {
	hotelId := ctx.Param("id")

	hotel, err := h.hotelService.GetHotelById(hotelId)
	if err != nil {
		response.WithError(ctx, http.StatusInternalServerError, messages.SomethingWentWrong, err)
		return
	}

	response.WithSuccess(ctx, http.StatusOK, "", gin.H{
		"hotel": hotel,
	})
}
