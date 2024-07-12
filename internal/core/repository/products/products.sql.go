package products

const (
	insertIntoProducts string = `
		INSERT INTO products (image, price, name, description, supplier_id, category_id, status, spec)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	updateProductImage string = `
		UPDATE products
		SET image = image || ',' || $1
		WHERE id = $2;
	`

	selectLimit string = `
		SELECT *
		FROM products
		LIMIT $1;
	`

	selectByID string = `
		SELECT *
		FROM products
		WHERE id = $1;
	`

	deleteByID = `
		DELETE FROM products
		WHERE id = $1;
	`

	updateByID = `
		UPDATE products
		SET 
			name = $1,
			price = $2,
			description = $3,
			supplier_id = $4,
			category_id = $5,
			spec = $6,
			satus = $7,
			created = now() AT TIME ZONE 'utc',
		WHERE id = $9;
	`
	searchByKeyword = `
		SELECT *
		FROM products
		WHERE
			name ILIKE '%$1%' or 
			description ILIKE '%$1%';
	
	`
)
