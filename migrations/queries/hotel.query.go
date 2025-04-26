package queries

const InsertHotelQuery = `
INSERT INTO hotels (
			id,
			name,
			description,
			city,
			country,
			image_url,
			price_per_night,
			rating,
			phone_number,
			features,
			created_at
		)
		VALUES (
			@id, @name, @description, @city, @country, @image_url, @price_per_night, @rating, @phone_number, @features, @created_at
		);
`
