package schemas

func All() []string {
	return []string{refreshTokens, users}
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
