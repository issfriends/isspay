package service

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type OrderDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error

	// Product
	UpdateProductQuantity(ctx context.Context, productID uint64, delta int64) error
	// Order
	CreateOrder(ctx context.Context, order *model.Order) error
	GetOrder(ctx context.Context, q *query.GetOrderQuery) error
	UpdateOrder(ctx context.Context, q *query.GetOrderQuery, updateOrder *model.Order) error
	// Wallet
	UpdateWalletAmount(ctx context.Context, walletID uint64, delta decimal.Decimal, isPay bool) (balance decimal.Decimal, err error)
}

type OrderServicer interface {
	CreateOrder(ctx context.Context, order *model.Order) (balance decimal.Decimal, err error)
	CancelOrder(ctx context.Context, orderID uint64) (balance decimal.Decimal, err error)
}

func NewOrder(db OrderDatabaser) OrderServicer {
	return &orderSvc{OrderDatabaser: db}
}

type orderSvc struct {
	OrderDatabaser
}

func (svc orderSvc) CreateOrder(ctx context.Context, order *model.Order) (decimal.Decimal, error) {
	var (
		balance decimal.Decimal
		err     error
	)

	err = svc.OrderDatabaser.ExecuteTx(ctx, func(txCtx context.Context) error {
		balance, err = svc.OrderDatabaser.UpdateWalletAmount(txCtx, uint64(order.WalletID), order.Amount.Neg(), false)
		if err != nil {
			return err
		}

		for _, op := range order.OrderedProducts {
			err = svc.OrderDatabaser.UpdateProductQuantity(txCtx, op.ProductID, -int64(op.Quantity))
			if err != nil {
				return err
			}
		}

		err = svc.OrderDatabaser.CreateOrder(txCtx, order)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return decimal.Zero, nil
	}

	return balance, nil
}

func (svc orderSvc) CancelOrder(ctx context.Context, orderID uint64) (decimal.Decimal, error) {
	var (
		balance decimal.Decimal
		err     error
	)

	err = svc.OrderDatabaser.ExecuteTx(ctx, func(txCtx context.Context) error {
		getOrderQ := &query.GetOrderQuery{
			ID:                 int64(orderID),
			HasOrderedProducts: true,
		}

		err = svc.OrderDatabaser.GetOrder(txCtx, getOrderQ)
		if err != nil {
			return err
		}
		order := getOrderQ.Data

		balance, err = svc.OrderDatabaser.UpdateWalletAmount(txCtx, uint64(order.WalletID), order.Amount, false)
		if err != nil {
			return err
		}

		for _, op := range order.OrderedProducts {
			err = svc.OrderDatabaser.UpdateProductQuantity(txCtx, op.ProductID, int64(op.Quantity))
			if err != nil {
				return err
			}
		}

		updateOrder := &model.Order{Status: value.Canceled}
		err = svc.OrderDatabaser.UpdateOrder(txCtx, getOrderQ, updateOrder)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return decimal.Zero, nil
	}

	return balance, nil
}
