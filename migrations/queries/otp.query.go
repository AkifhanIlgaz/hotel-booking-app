package queries

const InsertOTPToken = `
	INSERT INTO otp_tokens (id, email, token_hash, expires_at, created_at)
		VALUES (
			@id,
			@email,
			@token_hash,
			@expires_at,
			@created_at
		)
`

const SelectOTPToken = `
	SELECT id, email, token_hash, expires_at, created_at
	FROM otp_tokens
	WHERE email = @email AND token_hash = @token_hash
`

const DeleteOTPToken = `
	DELETE FROM otp_tokens
	WHERE token_hash = @token_hash
`

const SelectUserIdByEmail = `
	SELECT id
	FROM users
	WHERE email = @email
`
