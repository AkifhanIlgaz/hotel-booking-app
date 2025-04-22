package schemas

func All() []string {
	return []string{refreshTokens, users, otpTokens, hotels}
}

const refreshTokens string = `
CREATE TABLE IF NOT EXISTS refresh_tokens (
  	id UUID PRIMARY KEY,
  	user_id UUID NOT NULL UNIQUE REFERENCES users(id),
  	token_hash TEXT NOT NULL UNIQUE,
  	created_at TIMESTAMP NOT NULL,
  	expires_at TIMESTAMP NOT NULL
);
`

const users string = `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL CHECK (char_length(name) >= 3),
    email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    password_hash TEXT NOT NULL CHECK (char_length(password_hash) >= 8),
    role VARCHAR(10) NOT NULL CHECK (role IN ('admin', 'user')),
    created_at TIMESTAMP NOT NULL
);
`

const otpTokens string = `
CREATE TABLE IF NOT EXISTS otp_tokens (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);
`

const hotels string = `
CREATE TABLE IF NOT EXISTS hotels (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL CHECK (char_length(name) >= 3),
    description TEXT NOT NULL CHECK (char_length(description) >= 10),
    city VARCHAR(50) NOT NULL CHECK (char_length(city) >= 3),
    country VARCHAR(50) NOT NULL CHECK (char_length(country) >= 3),
    image_url VARCHAR(255) NOT NULL CHECK (char_length(image_url) >= 10),
    price_per_night NUMERIC(10, 2) NOT NULL CHECK (price_per_night > 0),
    rating NUMERIC(2, 1) NOT NULL CHECK (rating >= 0 AND rating <= 5),
    phone_number VARCHAR(20) NOT NULL,
    features TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
`
