package queries

const InsertOTPToken = `
	INSERT INTO otp_tokens (id, user_id, token, expires_at, created_at, used)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
`

const SelectOTPToken = `
	SELECT id, user_id, token, expires_at, created_at, used
	FROM otp_tokens
	WHERE user_id = $1 AND token = $2
`

const MarkOTPAsUsed = `
	UPDATE otp_tokens
	SET used = true
	WHERE id = $1
`

const SelectUserIdByEmail = `
	SELECT id
	FROM users
	WHERE email = $1
`
