package products

const (
	insertIntoProducts string = `
		INSERT INTO products (image, price, name, description, supplier_id, category_id, status, spec)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	updateProductImage string = `
		UPDATE products
		SET image = CASE
						WHEN image IS NULL OR image = '' THEN $1
						ELSE image || ',' || $1
					END
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
			name = CASE
						WHEN $1 <> '' THEN $1
						ELSE name 
					END,
			price = CASE
						WHEN $2 <> '' THEN $2
						ELSE price
					END,
			description = CASE
							WHEN $3 <> '' THEN $3
							ELSE description
						END,
			supplier_id = CASE
							WHEN $4 <> '' THEN $4
							ELSE supplier_id
						END,
			category_id = CASE
							WHEN $5 <> '' THEN $5
							ELSE category_id
						END,
			spec = CASE
						WHEN $6 <> '' THEN $6
						ELSE spec
					END,
			satus = CASE
						WHEN $7 <> '' THEN $7
						ELSE status
					END,
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
