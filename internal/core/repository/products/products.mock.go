package products

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ IProductRepository = (*Mock)(nil)

func NewProductsMock() *Mock {
	return &Mock{}
}

// GetLimit implements domain.IProductRepository.
func (p *Mock) GetLimit(ctx context.Context, limit int) ([]domain.ProductRes, error) {
	panic("unimplemented")
}

// Insert implements domain.IProductRepository.
func (p *Mock) Insert(ctx context.Context, prd *domain.Products) (int64, error) {
	panic("unimplemented")
}

// UploadNewImage implements domain.IProductRepository.
func (p *Mock) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	panic("unimplemented")
}
