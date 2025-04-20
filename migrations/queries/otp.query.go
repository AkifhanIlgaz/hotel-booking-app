package queries

const InsertOTPToken = `
	INSERT INTO otp_tokens (id, email, token_hash, expires_at, created_at)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
`

const SelectOTPToken = `
	SELECT id, email, token_hash, expires_at, created_at
	FROM otp_tokens
	WHERE email = $1 AND token_hash = $2
`

const DeleteOTPToken = `
	DELETE FROM otp_tokens
	WHERE token_hash = $1
`

const SelectUserIdByEmail = `
	SELECT id
	FROM users
	WHERE email = $1
`
