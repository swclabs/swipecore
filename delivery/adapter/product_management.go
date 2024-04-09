// product management adapter, base on base.adapt.go

package adapter

import "swclabs/swipe-api/internal/http"

const TypeProductManagement = "ProductManagement"

type _ProductManagementAdapter struct {
	server http.IServer
}

func _NewProductManagement() IAdapter {
	product := &_ProductManagementAdapter{
		server: http.New(),
	}
	return product
}

func (product *_ProductManagementAdapter) Run(addr string) error {
	product.server.Bootstrap(http.ProductManagementModule)
	return product.server.Run(addr)
}