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

	selectById string = `
		SELECT *
		FROM products
		WHERE id = $1;
	`

	deleteById = `
		DELETE FROM products
		WHERE id = $1;
	`
)
