package queries

const SelectUserByEmail = `
	SELECT *
	FROM users
	WHERE email = @email;
	`
