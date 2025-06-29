
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

// Hotels godoc
// @Summary      Get hotels
// @Description  get hotels
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        name query string false "name of the hotel"
// @Param        location query string false "location of the hotel"
// @Param        min_price query int false "minimum price of the hotel"
// @Param        max_price query int false "maximum price of the hotel"
// @Param        rating query int false "rating of the hotel"
// @Param        features query string false "features of the hotel"
// @Param        page query int false "page number"
// @Param        limit query int false "limit of the hotels"
// @Success      200  {object}  response.SuccessResponse{data=models.PaginatedHotels}
// @Failure      400  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /hotels [get]
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

// Hotel godoc
// @Summary      Get a hotel
// @Description  get a hotel by id
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Hotel ID"
// @Success      200  {object}  response.SuccessResponse{data=models.Hotel}
// @Failure      500  {object}  response.ErrorResponse
// @Router       /hotels/{id} [get]
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
