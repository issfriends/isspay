package bot

import (
	"github.com/google/uuid"
	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/pkg/chatbot"
)

// Ordering ordering domain handler
func (h Handler) Ordering() OrderingHandler {
	return OrderingHandler{svc: h.svc}
}

// OrderingHandler order and inventory handler
type OrderingHandler struct {
	svc *app.Service
}

// ListProductsEndpoint list products endpoint
func (h OrderingHandler) ListProductsEndpoint(c *chatbot.MsgContext) error {
	ctx := c.Ctx
	category := value.Drink
	if c.GetValue("category") == "snake" {
		category = value.Snake
	}

	q := &query.ListProductsQuery{
		Category:    category,
		QuantityGte: 1,
	}

	total, err := h.svc.Inventory.ListProducts(ctx, q)
	if err != nil {
		return err
	}

	round := total / 10
	if total%10 > 0 {
		round++
	}
	msgs := make([]chatbot.Message, 0, round)
	for i := 0; i < int(round); i++ {
		start := i * 10
		end := start + 10
		if end > int(total) {
			end = int(total)
		}
		msg := view.ProductsMenuMsg(q.Data[start:end])
		msgs = append(msgs, msg)
	}
	msgs = append(msgs, view.ShopMenuView("再來一點"))
	msgID := c.GetMessengerID()
	return c.PushMsgs(msgID, msgs...)
}

// PurchaseProductEndpoint purchase product endpoint
func (h OrderingHandler) PurchaseProductEndpoint(c *chatbot.MsgContext) error {
	var (
		orderedProduct = &model.OrderedProduct{}
		order          = &model.Order{}
		ctx            = c.Ctx
	)

	if err := c.Bind(orderedProduct, "json"); err != nil {
		return err
	}
	order.OrderedProducts = []*model.OrderedProduct{orderedProduct}
	order.UID = uuid.New().String()

	msgID := ""
	balance, err := h.svc.Ordering.CreateOrderByMsgID(ctx, msgID, order)
	if err != nil {
		return err
	}

	return c.ReplyMsg(
		chatbot.TextMsgf("購買成功，目前錢包餘額 %s", balance.StringFixed(2)),
		view.ShopMenuView("再來一點", view.OrderCancelBtn(order.UID)),
	)
}

// CancelOrderEndpoint cancel order endpoint
func (h OrderingHandler) CancelOrderEndpoint(c *chatbot.MsgContext) error {
	var (
		orderUID = c.GetValue("orderUID")
		ctx      = c.Ctx
	)

	msgID := ""
	balance, err := h.svc.Ordering.CancelOrderByMsgID(ctx, msgID, orderUID)
	if err != nil {
		return err
	}

	return c.ReplyMsg(chatbot.TextMsgf("訂單取消成功，目前錢包餘額 %s", balance.StringFixed(2)), view.ShopMenuView("再來一點"))
}
