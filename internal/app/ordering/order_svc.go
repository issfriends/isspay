package ordering

import (
	"context"
	"errors"
	"time"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type OrderDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, q *query.GetOrderQuery, updateOrder *model.Order) error
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
	GetOrder(ctx context.Context, q *query.GetOrderQuery) error
}

type AccountDatabaser interface {
	GetWallet(ctx context.Context, q *query.GetWalletQuery) error
	UpdateWalletAmount(ctx context.Context, walletID int64, delta decimal.Decimal, isPay bool) (balance decimal.Decimal, err error)
}

type InventoryDatabaser interface {
	UpdateProductQuantity(ctx context.Context, productID int64, delta int64) error
}
type OrderServicer interface {
	CreateOrder(ctx context.Context, walletQuery *query.GetWalletQuery, order *model.Order) (walletBalance decimal.Decimal, err error)
	CancelOrder(ctx context.Context, wallet *query.GetWalletQuery, orderUID string) (walletBalance decimal.Decimal, err error)
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
}

func (svc orderingSvc) CreateOrder(ctx context.Context, walletQuery *query.GetWalletQuery, order *model.Order) (walletBalance decimal.Decimal, err error) {
	err = svc.orderDB.ExecuteTx(ctx, func(txCtx context.Context) error {
		err = svc.accountDB.GetWallet(ctx, walletQuery)
		if err != nil {
			return err
		}
		wallet := walletQuery.Data

		if !wallet.CanPurchase(time.Now()) {
			return errors.New("")
		}

		for _, op := range order.OrderedProducts {
			err := svc.inventoryDB.UpdateProductQuantity(txCtx, op.ProductID, -int64(op.Quantity))
			if err != nil {
				return err
			}
		}

		order.Setup(wallet.ID)
		// 錢包扣錢
		walletBalance, err = svc.accountDB.UpdateWalletAmount(ctx, wallet.ID, order.Amount.Neg(), false)
		if err != nil {
			return err
		}

		// 建立訂單
		err = svc.orderDB.CreateOrder(txCtx, order)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return walletBalance, err
	}

	return walletBalance, nil
}

func (svc orderingSvc) CancelOrder(ctx context.Context, walletQuery *query.GetWalletQuery, orderUID string) (walletBalance decimal.Decimal, err error) {
	err = svc.orderDB.ExecuteTx(ctx, func(txCtx context.Context) error {
		err = svc.accountDB.GetWallet(ctx, walletQuery)
		if err != nil {
			return err
		}
		wallet := walletQuery.Data
		orderQuery := &query.GetOrderQuery{UID: orderUID, HasOrderedProducts: true}
		if err = svc.orderDB.GetOrder(ctx, orderQuery); err != nil {
			return err
		}
		order := orderQuery.Data

		// recover wallet balance
		walletBalance, err = svc.accountDB.UpdateWalletAmount(ctx, wallet.ID, order.Amount, false)
		if err != nil {
			return err
		}
		// recovert product quantity
		for _, op := range order.OrderedProducts {
			err = svc.inventoryDB.UpdateProductQuantity(ctx, op.ProductID, int64(op.Quantity))
			if err != nil {
				return err
			}
		}

		err = svc.orderDB.UpdateOrder(ctx, orderQuery, &model.Order{Status: value.Canceled})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return walletBalance, err
	}

	return walletBalance, nil
}

func (svc orderingSvc) ListOrders(ctx context.Context, q *query.ListOrdersQuery) error {

	return nil
}

func (svc orderingSvc) GetOrder(ctx context.Context, q *query.GetOrderQuery) error {
	// 鄭言竹交給你了
	if q == nil {
		q = &query.GetOrderQuery{}
	}
	err := svc.orderDB.GetOrder(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
