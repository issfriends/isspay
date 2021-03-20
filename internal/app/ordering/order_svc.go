package ordering

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model/value"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type OrderDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context) error
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
	GetOrder(ctx context.Context, q *query.GetOrderQuery) error
}

type AccountDatabaser interface {
	UpdateWallet(ctx context.Context, q *query.GetWalletQuery, wallet *model.Wallet) error
}

type OrderServicer interface {
	CreateOrderByMsgID(ctx context.Context, msgID string, order *model.Order) (walletBalance decimal.Decimal, err error)
	CancelOrderByMsgID(ctx context.Context, msgID string, orderUID string) (walletBalance decimal.Decimal, err error)
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
}

func (svc orderingSvc) CreateOrderByMsgID(ctx context.Context, msgID string, order *model.Order) (walletBalance decimal.Decimal, err error) {
	// 鄭言竹交給你了
	var (
		orderAmount  decimal.Decimal
		updateWallet = &model.Wallet{}
	)

	for _, op := range order.OrderedProducts {
		price := op.Price.Mul(decimal.NewFromInt(op.Quantity))
		orderAmount = orderAmount.Add(price)
		op.OrderID = order.ID
	}

	err = svc.orderDB.ExecuteTx(ctx, func(txCtx context.Context) error {
		// 錢包扣錢
		whereWallet := &query.GetWalletQuery{MessengerID: msgID}
		err = svc.accountDB.UpdateWallet(txCtx, whereWallet, updateWallet)
		if err != nil {
			return err
		}
		// 建立訂單
		err = svc.orderDB.CreateOrder(txCtx, &model.Order{
			UID:             order.UID,
			WalletID:        order.WalletID,
			Status:          value.Completed,
			Amount:          orderAmount,
			OrderedProducts: order.OrderedProducts,
		})
		if err != nil {
			return err
		}
		return nil
	})

	// TODO 計算餘額...
	return decimal.Zero, nil
}

func (svc orderingSvc) CancelOrderByMsgID(ctx context.Context, msgID string, orderUID string) (walletBalance decimal.Decimal, err error) {
	// 鄭言竹交給你了
	return decimal.Zero, nil
}

func (svc orderingSvc) ListOrders(ctx context.Context, q *query.ListOrdersQuery) error {
	// 鄭言竹交給你了
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
