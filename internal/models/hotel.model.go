package models

import (
	"math"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Hotel struct {
	Id            uuid.UUID `json:"id" `
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Location      Location  `json:"location"`
	ImageUrl      string    `json:"image_url"`
	PricePerNight string    `json:"price_per_night"`
	PhoneNumber   string    `json:"phone_number"`
	Rating        float64   `json:"rating"`
	Features      []string  `json:"features"`
	CreatedAt     string    `json:"created_at"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type HotelFilterParams struct {
	Page      int      `json:"page" form:"page"`
	PageSize  int      `json:"pageSize" form:"pageSize" `
	SortBy    string   `json:"sortBy" form:"sortBy" `
	SortOrder string   `json:"sortOrder" form:"sortOrder"`
	City      string   `json:"city" form:"city" `
	Country   string   `json:"country" form:"country"`
	MinPrice  int      `json:"minPrice" form:"minPrice"`
	MaxPrice  int      `json:"maxPrice" form:"maxPrice"`
	MinRating float64  `json:"minRating" form:"minRating"`
	Features  []string `json:"features" form:"features"`
	Search    string   `json:"search" form:"search"`
}

// TODO: Create separate function to set defaults ?
// TODO: return error
func (p *HotelFilterParams) Validate() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.PageSize <= 0 || p.PageSize >= 50 {
		p.PageSize = 5
	}

	if p.SortBy == "" {
		p.SortBy = "name"
	}

	if p.MinPrice < 0 {
		p.MinPrice = 0
	}

	if p.MaxPrice == 0 {
		p.MaxPrice = math.MaxInt
	}
}

func (p *HotelFilterParams) NormalizeFeatures() {
	concatenatedFeatures := strings.Join(slices.Concat(p.Features), ",")
	p.Features = strings.Split(concatenatedFeatures, ",")

	uniqueFeatures := make(map[string]struct{})
	for _, f := range p.Features {
		cleaned := strings.ToLower(strings.TrimSpace(f))
		if cleaned != "" {
			uniqueFeatures[cleaned] = struct{}{}
		}
	}

	p.Features = make([]string, 0, len(uniqueFeatures))
	for feature := range uniqueFeatures {
		p.Features = append(p.Features, feature)
	}
}
