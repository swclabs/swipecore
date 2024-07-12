// Package addresses  Duc Hung Ho @kyeranyo
package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/infra/db"
)

// Addresses struct for address repository
type Addresses struct {
	db db.IDatabase
}

// New creates a new Addresses object
func New(conn db.IDatabase) IAddressRepository {
	return useCache(&Addresses{
		db: conn,
	})
}

// Insert implements domain.IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data domain.Addresses) error {
	return addr.db.SafeWrite(
		ctx, insertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.UUID,
	)
}
