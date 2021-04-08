package bot

import (
	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/shopspring/decimal"
)

func (h Handler) Account() AccountHandler {
	return AccountHandler(h)
}

type AccountHandler struct {
	svc *app.App
}

func (h AccountHandler) SignUpEndpoint(c *chatbot.MsgContext) error {
	account := &model.Account{}
	if err := c.Bind(account, "json"); err != nil {
		return err
	}
	userProfile, err := c.GetCurrentUser()
	if err != nil {
		return err
	}
	account.UserName = userProfile.Username
	if err := account.MessengerID.Scan(c.GetMessengerID()); err != nil {
		return err
	}

	err = h.svc.Auth.SignUpByChatbot(c.Ctx, account)
	if err != nil {
		return err
	}

	return c.ReplyTextf("welcome %s", account.NickName)
}

func (h AccountHandler) SwitchMemberEndpoint(c *chatbot.MsgContext) error {
	email := c.GetValue("membership")
	return c.ReplyTextf("welcome %s", email)
}

func (h AccountHandler) PaymentEndpoint(c *chatbot.MsgContext) error {
	var (
		ctx = c.Ctx
	)
	amount, err := decimal.NewFromString(c.GetValue("amount"))
	if err != nil {
		return err
	}

	balance, err := h.svc.Account.MakePayment(ctx, 0, amount)
	if err != nil {
		return err
	}

	return c.ReplyTextf("還款成功，目前錢包餘額 %s", balance.StringFixed(2))
}

func (h AccountHandler) GetBalanceEndpoint(c *chatbot.MsgContext) error {
	var (
		q = &query.GetWalletQuery{
			MessengerID: c.GetMessengerID(),
		}
		ctx = c.Ctx
	)

	if err := h.svc.Account.GetWallet(ctx, q); err != nil {
		return err
	}

	msg := view.GetWalletView(q.Data)
	return c.ReplyMsg(msg)
}
