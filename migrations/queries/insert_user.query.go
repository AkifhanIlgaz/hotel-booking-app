package queries

const InsertUser = `
		INSERT INTO users (
			id,
			name,
			email,
			password_hash,
			role,
			created_at
		)
		VALUES (
			@id, @name, @email, @password_hash, @role, @created_at
		);
`

const UpdateUserPasswordByEmail = `
	UPDATE users
	SET password_hash = @password_hash
	WHERE email = @email;
`
