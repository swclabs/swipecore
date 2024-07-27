package inventories

const (
	insertIntoInventory string = `
		INSERT INTO inventories (product_id, price, specs, available, currency_code, status, image, color, color_img)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	getByID string = `
		SELECT * FROM inventories WHERE id = $1;
	`

	getByProductID = `
		SELECT * FROM inventories WHERE product_id = $1;
	`

	getProductsLimit = `
		SELECT * FROM inventories LIMIT $1 OFFSET $2;
	`

	deleteInventorybyID = `
		DELETE FROM inventories WHERE id = $1;
	`

	uploadInventoryImage = `
		UPDATE inventories
		SET image = CASE
						WHEN image is NULL OR image = '' THEN $1
						ELSE image || ',' || $1
					END
		)
		WHERE id = $2;
	`

	update = `
		UPDATE inventories
		SET product_id = CASE
							WHEN $2 <> 0 THEN $2
							ELSE product_id
						END,
			status = CASE 
						WHEN $3 <> '' THEN $3 
						ELSE status 
					END,
			price = CASE
						WHEN $4 <> '' THEN $4
						ELSE price
					END,
			currency_code = CASE
								WHEN $5 <> '' THEN $5
								ELSE currency_code
							END,
			specs = CASE
						WHEN $6 <> '' THEN $6
						ELSE specs
					END,
			available = CASE
							WHEN $7 <> '' THEN $7
							ELSE available
						END,
			image = CASE
						WHEN $8 <> '' THEN $8
						ELSE image
					END,
			color = CASE
						WHEN $9 <> '' THEN $9
						ELSE color
					END,
			color_img = CASE
							WHEN $10 <> '' THEN $10
							ELSE color_img
						END
		WHERE id = $1;
	`
)
