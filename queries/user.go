package queries

const (
	UserGet = `
		SELECT id
		FROM users
		WHERE email = $1
	`

	UserCreate = `
		INSERT INTO users(name,email)
			VALUES($1,$2)
		RETURNING id;
	`

	UpdateLastLogin = `
		UPDATE users
			SET last_login=CURRENT_TIMESTAMP
		WHERE id=$1;
	`
)
