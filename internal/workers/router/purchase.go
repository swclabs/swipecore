// Package router define tasks - queue
package router

import (
	"swclabs/swix/internal/workers/handler"
	"swclabs/swix/internal/workers/server"
)

// IPurchase interface for Purchase objects
type IPurchase interface {
	IRouter
}

// NewPurchase creates a new Purchase object
func NewPurchase(handler handler.IPurchase) IPurchase {
	return &Purchase{
		handler: handler,
	}
}

// Purchase struct define the Purchase object
type Purchase struct {
	handler handler.IPurchase
}

// Register implements IPurchase.
func (p *Purchase) Register(eng server.IWorker) {
	eng.RegisterQueue(p.handler.HandleAddToCart)
}
