package queries

const (
	AddressGet = `
		SELECT id,
			   user_id,
			   address,
			   title,
			   created
		FROM addresses
		WHERE id = $1 AND user_id = $2
	`

	AddressUser = `
		SELECT id,
			   	user_id,
				address,
				title,
				created
		FROM addresses
		WHERE user_id = $1
		ORDER BY created ASC;
	`

	AddressUserCount = `
		SELECT COUNT(*) FROM addresses
		WHERE user_id=$1
	`

	AddressCreate = `
		INSERT INTO addresses(user_id,address,title)
		VALUES($1,$2,$3)
		RETURNING id
	`

	AddressUpdate = `
		UPDATE addresses SET address=$1, title=$2
		WHERE id=$3 AND user_id=$4
	`

	AddressDelete = `
		DELETE FROM addresses
		WHERE id=$1 AND user_id=$2
	`
)
