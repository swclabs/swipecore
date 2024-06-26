package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ISuppliersRepository implements all methods to access and manage supplier data in the database.
type ISuppliersRepository interface {
	// Insert adds a new supplier and their address to the database.
	// ctx is the context to manage the request's lifecycle.
	// sup is the Suppliers object to be added.
	// addr is the Addresses object associated with the supplier.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, sup domain.Suppliers, addr domain.Addresses) error

	// InsertAddress adds a new address for a supplier to the database.
	// ctx is the context to manage the request's lifecycle.
	// addr is the SuppliersAddress object to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertAddress(ctx context.Context, addr domain.SuppliersAddress) error

	// GetLimit retrieves a list of suppliers with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of suppliers to retrieve.
	// Returns a slice of Suppliers objects and an error if any issues occur during the retrieval process.
	GetLimit(ctx context.Context, limit int) ([]domain.Suppliers, error)

	// GetByPhone retrieves a supplier by their phone number.
	// ctx is the context to manage the request's lifecycle.
	// phone is the phone number to search for.
	// Returns a pointer to the Suppliers object and an error if any issues occur during the retrieval process.
	GetByPhone(ctx context.Context, phone string) (*domain.Suppliers, error)
}
