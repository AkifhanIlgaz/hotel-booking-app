package queries

const InsertRefreshToken string = `
  INSERT INTO refresh_tokens (id, user_id, hashed_token, created_at, expires_at)
    VALUES ($1, $2, $3, $4, $5)
`
