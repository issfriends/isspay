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
	*app.App
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

	err = h.Auth.SignUpByChatbot(c.Ctx, account)
	if err != nil {
		return err
	}

	return Replyi18nText(c, "sign_up_reply", account)
}

func (h AccountHandler) SwitchMemberEndpoint(c *chatbot.MsgContext) error {
	var (
		account = &model.Account{}
		getAccQ = &query.GetAccountQuery{}
		ctx     = c.Ctx
	)

	err := account.Role.FromString(c.GetValue("membership"))
	if err != nil {
		return err
	}

	claims, err := GetClaims(c)
	if err != nil {
		return err
	}

	getAccQ.ID = int64(claims.AccountID)

	err = h.Account.UpdateAccount(ctx, getAccQ, account)
	if err != nil {
		return err
	}

	return Replyi18nText(c, "switch_membership_reply", account)
}

func (h AccountHandler) PaymentEndpoint(c *chatbot.MsgContext) error {
	var (
		ctx = c.Ctx
	)
	amount, err := decimal.NewFromString(c.GetValue("amount"))
	if err != nil {
		return err
	}

	balance, err := h.Account.MakePayment(ctx, 0, amount)
	if err != nil {
		return err
	}

	return Replyi18nText(c, "payment_reply", map[string]interface{}{
		"Amount":  amount,
		"Balance": balance,
	})
}

func (h AccountHandler) GetBalanceEndpoint(c *chatbot.MsgContext) error {
	var (
		q = &query.GetWalletQuery{
			MessengerID: c.GetMessengerID(),
		}
		ctx = c.Ctx
	)

	if err := h.Account.GetWallet(ctx, q); err != nil {
		return err
	}

	msg, err := view.GetWalletView(q.Data)
	if err != nil {
		return err
	}
	return c.ReplyMsg(msg)
}
