package queries

// TODO: Upsert
const InsertRefreshToken string = `
INSERT INTO refresh_tokens (id, user_id, token_hash, created_at, expires_at)
	VALUES (@id, @user_id, @token_hash, @created_at, @expires_at)
`

const ExpiryCheck string = `
SELECT expires_at FROM refresh_tokens
WHERE user_id = @user_id;
`

const UpdateRefreshToken string = `
UPDATE refresh_tokens
SET token_hash = @token_hash, 
    created_at = @created_at,
    expires_at = @expires_at
    WHERE user_id = @user_id;
`

const SelectRefreshToken string = `
SELECT * FROM refresh_tokens
WHERE token_hash = @token_hash;
`

const DeleteRefreshToken string = `
DELETE FROM refresh_tokens
WHERE user_id = @user_id;
`
