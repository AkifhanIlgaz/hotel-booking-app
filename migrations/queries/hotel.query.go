package queries

const InsertHotelQuery = `
INSERT INTO hotels (id, name, description, city, country, image_url, price_per_night, rating, phone_number, features, created_at)	
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`
