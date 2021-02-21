package bot

import (
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/issfriends/isspay/pkg/factory"
)

func (h Handler) Ordering() OrderingHandler {
	return OrderingHandler{}
}

type OrderingHandler struct {
}

func (h OrderingHandler) ListProductsEndpoint(c *chatbot.MsgContext) error {

	products := factory.Product.MustBuildN(10).([]*model.Product)
	msg1 := view.ProductsMenuMsg(products)

	products = factory.Product.MustBuildN(10).([]*model.Product)
	msg2 := view.ProductsMenuMsg(products)
	return c.ReplyMsg(msg1, msg2, view.ShopMenuView("再來一點"))
}

func (h OrderingHandler) PurchaseProductEndpoint(c *chatbot.MsgContext) error {
	productUID := c.GetValue("productUID")

	return c.ReplyMsg(chatbot.TextMsgf("買了 %s", productUID), view.ShopMenuView("再來一點", view.OrderCancelBtn(productUID)))
}

func (h OrderingHandler) CancelOrderEndpoint(c *chatbot.MsgContext) error {
	orderUID := c.GetValue("orderUID")
	return c.ReplyMsg(chatbot.TextMsgf("order (%s) 取消了", orderUID), view.ShopMenuView("再來一點"))
}
