package models

// TODO: Status enum  => pending, paid

type Reservation struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	HotelId      string `json:"hotel_id"`
	CheckInDate  string `json:"check_in_date"`
	CheckOutDate string `json:"check_out_date"`
	GuestCount   int    `json:"guest_count"`
	TotalPrice   int    `json:"total_price"`
	Status       string `json:"status"`
}
