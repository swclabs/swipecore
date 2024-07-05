package inventories

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/errors"
	"swclabs/swipecore/pkg/db"
)

type Inventory struct {
	db db.IDatabase
}

var _ IInventoryRepository = (*Inventory)(nil)

func New(conn db.IDatabase) IInventoryRepository {
	return useCache(&Inventory{
		db: conn,
	})
}

// GetLimit implements IInventoryRepository.
func (w *Inventory) GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error) {
	if offset <= 1 {
		offset = 1
	}
	rows, err := w.db.Query(ctx, getProductsLimit, limit, (offset-1)*limit)
	if err != nil {
		return nil, err
	}
	return db.CollectRows[domain.Inventories](rows)
}

func (w *Inventory) GetByProductId(ctx context.Context, productId int64) ([]domain.Inventories, error) {
	rows, err := w.db.Query(ctx, getByProductId, productId)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	inventories, err := db.CollectRows[domain.Inventories](rows)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	return inventories, nil
}

// GetById implements IInventoryRepository.
func (w *Inventory) GetById(ctx context.Context, inventoryId int64) (*domain.Inventories, error) {
	rows, err := w.db.Query(ctx, getById, inventoryId)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[domain.Inventories](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// FindDevice implements domain.IInventoryRepository.
func (w *Inventory) FindDevice(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.Inventories, error) {
	rows, err := w.db.Query(ctx, getAvailableProducts,
		deviceSpecs.ProductId, deviceSpecs.Ram, deviceSpecs.Ssd, deviceSpecs.Color)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[domain.Inventories](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// InsertProduct implements domain.IInventoryRepository.
func (w *Inventory) InsertProduct(
	ctx context.Context, product domain.InventoryStruct) error {
	specsjson, _ := json.Marshal(product.Specs)
	return w.db.SafeWrite(ctx, insertIntoInventory,
		product.ProductID, product.Model, product.Price,
		string(specsjson), product.Available, product.CurrencyCode,
		"active",
	)
}