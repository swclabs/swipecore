package domain

const ProductsTable = "products"

// Products Table
type Products struct {
	ID          int64  `json:"id" gorm:"column:id"`
	Image       string `json:"image" gorm:"column:image"`
	Price       string `json:"price" gorm:"column:price"`
	Description string `json:"description" gorm:"column:description"`
	Name        string `json:"name" gorm:"column:name"`
	SupplierID  string `json:"supplier_id" gorm:"column:supplier_id"`
	CategoryID  string `json:"category_id" gorm:"column:category_id"`
	Available   string `json:"available" gorm:"column:available"`
	Spec        string `json:"spec" gorm:"column:spec"`
}

type ProductRequest struct {
	Price       string `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
	SupplierID  string `json:"supplierID" validate:"required"`
	CategoryID  string `json:"categoryID" validate:"required"`
	Available   string `json:"available" validate:"required"`
}

// ProductInCart Table
type ProductInCart struct {
	ID        int64 `json:"id" gorm:"column:id"`
	CartID    int64 `json:"cart_id" gorm:"column:cart_id"`
	ProductID int64 `json:"product_id" gorm:"column:product_id"`
	Amount    int64 `json:"amount" gorm:"column:amount"`
}

// ProductInOrder Table
type ProductInOrder struct {
	ID        int64 `json:"id" gorm:"column:id"`
	OrderID   int64 `json:"order_id" gorm:"column:order_id"`
	ProductID int64 `json:"product_id" gorm:"column:product_id"`
	Amount    int64 `json:"amount" gorm:"column:amount"`
}

// FavoriteProduct Table
type FavoriteProduct struct {
	ID        int64 `json:"id" gorm:"column:id"`
	UserID    int64 `json:"user_id" gorm:"column:user_id"`
	ProductID int64 `json:"product_id" gorm:"column:product_id"`
}

type Accessory struct {
	Name  string `json:"name" gorm:"column:name"`
	Price string `json:"price" gorm:"column:price"`
	New   string `json:"new" gorm:"column:new"`
	Img   string `json:"img" gorm:"column:img"`
}

type Specifications struct {
	Screen  string `json:"screen"`
	Display string `json:"display"`
	SSD     []int  `json:"SSD"`
}

type ProductResponse struct {
	ID          int64          `json:"id"`
	Image       string         `json:"image"`
	Price       string         `json:"price"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	Available   string         `json:"available"`
	Spec        Specifications `json:"spec"`
}

type ProductsListResponse struct {
	Data []Products `json:"data" gorm:"column:data"`
}
