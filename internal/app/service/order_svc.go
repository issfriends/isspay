package service

import "context"

type OrderDatabaser interface {
	// Product
	UpdateProductQuantity(ctx context.Context)
	// Order
	CreateOrder(ctx context.Context)
	// Wallet
	UpdateWalletAmount(ctx context.Context)
}

type OrderServicer interface {
	CreateOrder(ctx context.Context)
	CancelOrder(ctx context.Context)
}
