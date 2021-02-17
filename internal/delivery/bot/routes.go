package bot

import (
	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/internal/pkg/chatbot"
)

func New(svc *app.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

type Handler struct {
	svc *app.Service
}

func Routes(bot chatbot.ChatBot, h *Handler) {
	bot.SetCommand(view.MenuCmd, h.Menu().MenuEndpoint)

	bot.SetCommand(view.ListProductsCmd, h.Ordering().ListProductsEndpoint)
	bot.SetCommand(view.PurchaseProductCmd, h.Ordering().PurchaseProductEndpoint)
	bot.SetCommand(view.CancelOrderCmd, h.Ordering().CancelOrderEndpoint)

	bot.SetForm(view.SignUpForm.SetHandle(h.Account().SignUpEndpoint))
	bot.SetCommand(view.SwitchMemberCmd, h.Account().SwitchMemberEndpoint)

	bot.SetCommand(view.CheckBalanceCmd, h.Account().GetBalanceEndpoint)
	bot.SetCommand(view.PaymentCmd, h.Account().PaymentEndpoint)
}
