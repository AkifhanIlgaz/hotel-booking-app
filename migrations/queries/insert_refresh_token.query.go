package queries

// TODO: Upsert
const InsertRefreshToken string = `
INSERT INTO refresh_tokens (id, user_id, token_hash, created_at, expires_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id) DO UPDATE
SET token_hash = EXCLUDED.token_hash,
    expires_at = EXCLUDED.expires_at,
    created_at = EXCLUDED.created_at;
`
