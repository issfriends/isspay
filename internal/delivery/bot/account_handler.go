package bot

import (
	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/pkg/chatbot"
)

func (h Handler) Account() AccountHandler {
	return AccountHandler{
		svc: h.svc,
	}
}

type AccountHandler struct {
	svc *app.Service
}

func (h AccountHandler) SignUpEndpoint(c *chatbot.MsgContext) error {
	account := &model.Account{}
	if err := c.Bind(account, "json"); err != nil {
		return err
	}
	if err := account.MessengerID.Scan(c.GetMessengerID()); err != nil {
		return err
	}
	if err := h.svc.Account.SignUpByChatbot(c.Ctx, account); err != nil {
		return err
	}

	return c.PushTextf("welcome %s", account.NickName)
}

func (h AccountHandler) SwitchMemberEndpoint(c *chatbot.MsgContext) error {
	email := c.GetValue("membership")
	return c.PushTextf("welcome %s", email)
}

func (h AccountHandler) PaymentEndpoint(c *chatbot.MsgContext) error {
	return c.PushTextf("payment!! balance:0")
}

func (h AccountHandler) GetBalanceEndpoint(c *chatbot.MsgContext) error {
	return c.PushTextf("payment!! balance:-100")
}
