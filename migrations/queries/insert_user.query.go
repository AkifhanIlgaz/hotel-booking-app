package queries

const InsertUser = `
	INSERT INTO users (id, name, email, password_hash, role, created_at)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
`
