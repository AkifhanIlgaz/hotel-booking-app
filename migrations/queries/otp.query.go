package queries

const InsertOTPToken = `
	INSERT INTO otp_tokens (id, user_id, token, expires_at, created_at)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)
`

const SelectOTPToken = `
	SELECT id, user_id, token, expires_at, created_at
	FROM otp_tokens
	WHERE user_id = $1 AND token = $2
`

const DeleteOTPToken = `
	DELETE FROM otp_tokens
	WHERE token = $1
`

const SelectUserIdByEmail = `
	SELECT id
	FROM users
	WHERE email = $1
`
