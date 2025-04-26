package main

import (
	"fmt"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/queries"
)

func main() {
	fmt.Println(queries.BuildHotelsQueryWithParams(models.HotelFilterParams{
		Page:      1,
		PageSize:  3,
		SortBy:    "name",
		City:      "c",
		Country:   "G",
		MinPrice:  0,
		MaxPrice:  500,
		MinRating: 3,
		Features:  []string{"Private beach", "zoz"},
		Search:    "",
	}))
}
