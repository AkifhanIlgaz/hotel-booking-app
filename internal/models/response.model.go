package models

//lint:ignore U1000 we want to have this type in swagger
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//lint:ignore U1000 we want to have this type in swagger
type ResetToken struct {
	ResetToken string `json:"reset_token"`
}

//lint:ignore U1000 we want to have this type in swagger
type PaginatedHotels struct {
	Hotels []Hotel `json:"hotels"`
}