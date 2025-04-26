package schemas

func All() []string {
	return []string{users, refreshTokens, otpTokens, hotels}
}

const refreshTokens string = `
IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='refresh_tokens' AND xtype='U')
BEGIN
    CREATE TABLE refresh_tokens (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        user_id UNIQUEIDENTIFIER NOT NULL UNIQUE,
        token_hash NVARCHAR(255) NOT NULL UNIQUE,
        created_at DATETIME2 NOT NULL,
        expires_at DATETIME2 NOT NULL,
        CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES users(id)
    );
END

`

const users string = `
   IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='users' AND xtype='U')
BEGIN
    CREATE TABLE users (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        name NVARCHAR(50) NOT NULL,
        email NVARCHAR(255) NOT NULL UNIQUE,
        password_hash NVARCHAR(MAX) NOT NULL,
        role NVARCHAR(10) NOT NULL,
        created_at DATETIME2 NOT NULL,

        CONSTRAINT CHK_name_length CHECK (LEN(name) >= 3),
        CONSTRAINT CHK_password_length CHECK (LEN(password_hash) >= 8),
        CONSTRAINT CHK_role CHECK (role IN ('admin', 'user')),
        CONSTRAINT CHK_email_format CHECK (
            email LIKE '[A-Za-z0-9._%+-]%@[A-Za-z0-9.-]%.[A-Za-z][A-Za-z]%'
        )
    );
END

`

const otpTokens string = `
IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='otp_tokens' AND xtype='U')
BEGIN
    CREATE TABLE otp_tokens (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        email NVARCHAR(255) NOT NULL UNIQUE,
        token_hash NVARCHAR(MAX) NOT NULL,
        expires_at DATETIME2 NOT NULL,
        created_at DATETIME2 NOT NULL,

         CHECK (
            email LIKE '[A-Za-z0-9._%+-]%@[A-Za-z0-9.-]%.[A-Za-z][A-Za-z]%'
        )
    );
END

`

const hotels string = `
DROP TABLE IF EXISTS hotels;

IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='hotels' AND xtype='U')
BEGIN
    CREATE TABLE hotels (
        id UNIQUEIDENTIFIER PRIMARY KEY,
        name NVARCHAR(255) NOT NULL UNIQUE,
        description NVARCHAR(255) NOT NULL,
        city NVARCHAR(50) NOT NULL,
        country NVARCHAR(50) NOT NULL,
        image_url NVARCHAR(255) NOT NULL,
        price_per_night DECIMAL(10, 2) NOT NULL,
        rating DECIMAL(2, 1) NOT NULL,
        phone_number NVARCHAR(20) NOT NULL,
        features TEXT NOT NULL,
        created_at DATETIME2 NOT NULL,

        CONSTRAINT CHK_hotel_name_length CHECK (LEN(name) >= 3),
        CONSTRAINT CHK_hotel_description_length CHECK (LEN(description) >= 10),
        CONSTRAINT CHK_hotel_city_length CHECK (LEN(city) >= 3),
        CONSTRAINT CHK_hotel_country_length CHECK (LEN(country) >= 3),
        CONSTRAINT CHK_hotel_image_url_length CHECK (LEN(image_url) >= 10),
        CONSTRAINT CHK_hotel_price_positive CHECK (price_per_night > 0),
        CONSTRAINT CHK_hotel_rating_range CHECK (rating >= 0 AND rating <= 5)
    );
END

`
