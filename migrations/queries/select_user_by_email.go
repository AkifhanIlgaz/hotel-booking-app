package queries

const SelectUserByEmail = `
	SELECT *
	FROM users
	WHERE email = $1;
	`
