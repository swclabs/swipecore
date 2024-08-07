package categories

const (
	insertIntoCategory = `
		INSERT INTO categories (name, description) 
		VALUES ($1, $2)
	`

	selectCategoryLimit string = `
		SELECT *
		FROM categories
		LIMIT $1;
	`

	selectCategoryByID string = `
		SELECT *
		FROM categories
		WHERE id = $1;
	`
)
