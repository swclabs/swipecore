package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain/dto"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/db"
	"swclabs/swipecore/pkg/lib/errors"
)

var _ IInventoryRepository = (*Inventory)(nil)

// New creates a new Inventory object.
func New(conn db.IDatabase) IInventoryRepository {
	return useCache(&Inventory{
		db: conn,
	})
}

// Inventory represents the Inventory object.
type Inventory struct {
	db db.IDatabase
}

// Update implements IInventoryRepository.
func (w *Inventory) Update(ctx context.Context, inventory entity.Inventories) error {
	panic("unimplemented")
}

// UploadImage implements IInventoryRepository.
func (w *Inventory) UploadImage(ctx context.Context, ID int, url string) error {
	return w.db.SafeWrite(ctx, uploadInventoryImage, url, ID)
}

// DeleteByID implements IInventoryRepository.
func (w *Inventory) DeleteByID(ctx context.Context, inventoryID int64) error {
	return w.db.SafeWrite(ctx, deleteInventorybyID, inventoryID)
}

// GetLimit implements IInventoryRepository.
func (w *Inventory) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventories, error) {
	if offset < 1 {
		offset = 1
	}
	rows, err := w.db.Query(ctx, getProductsLimit, limit, (offset-1)*limit)
	if err != nil {
		return nil, err
	}
	return db.CollectRows[entity.Inventories](rows)
}

// GetByProductID implements IInventoryRepository.
func (w *Inventory) GetByProductID(ctx context.Context, productID int64) ([]entity.Inventories, error) {
	rows, err := w.db.Query(ctx, getByProductID, productID)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	inventories, err := db.CollectRows[entity.Inventories](rows)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	return inventories, nil
}

// GetByID implements IInventoryRepository.
func (w *Inventory) GetByID(ctx context.Context, inventoryID int64) (*entity.Inventories, error) {
	rows, err := w.db.Query(ctx, getByID, inventoryID)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[entity.Inventories](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// FindDevice implements IInventoryRepository.
func (w *Inventory) FindDevice(ctx context.Context, device dto.InventoryDeviceSpecs) (*entity.Inventories, error) {
	rows, err := w.db.Query(ctx, getAvailableProducts,
		device.ProductID, device.RAM, device.Ssd, device.Color)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[entity.Inventories](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// InsertProduct implements IInventoryRepository.
func (w *Inventory) InsertProduct(ctx context.Context, product entity.Inventories) error {
	return w.db.SafeWrite(ctx, insertIntoInventory,
		product.ProductID, product.Price,
		product.Specs, product.Available, product.CurrencyCode,
		product.Status,
	)
}
