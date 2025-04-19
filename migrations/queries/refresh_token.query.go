package queries

// TODO: Upsert
const InsertRefreshToken string = `
INSERT INTO refresh_tokens (id, user_id, token_hash, created_at, expires_at)
	VALUES ($1, $2, $3, $4, $5)
`

const ExpiryCheck string = `
SELECT expires_at FROM refresh_tokens
WHERE user_id = $1;
`

const UpdateRefreshToken string = `
UPDATE refresh_tokens
SET token_hash = $1, 
    created_at = $2,
    expires_at = $3
    WHERE user_id = $4;
`

const SelectRefreshToken string = `
SELECT * FROM refresh_tokens
WHERE token_hash = $1;
`
