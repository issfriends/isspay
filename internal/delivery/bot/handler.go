package bot

import (
	"path/filepath"
	"runtime"

	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/pkg/chatbot"
)

type (
	ClaimsKey struct{}
)

// New new a chatbot handler
func New(svc *app.App) *Handler {
	return &Handler{
		App: svc,
	}
}

// Handler chatbot handler
type Handler struct {
	*app.App
}

// Routes set chatbot routes
func (h *Handler) Routes(bot chatbot.ChatBot) error {
	if err := bot.SetMenu(view.DefaultMenu, assets("linebot_menu_v1")); err != nil {
		return err
	}

	bot.SetCommand(view.MenuCmd, h.Menu().MenuEndpoint)

	bot.SetCommand(view.ListProductsCmd, h.Ordering().ListProductsEndpoint, h.CheckAuth)
	bot.SetCommand(view.PurchaseProductCmd, h.Ordering().PurchaseProductEndpoint, h.CheckAuth)
	bot.SetCommand(view.CancelOrderCmd, h.Ordering().CancelOrderEndpoint, h.CheckAuth)

	bot.SetForm(view.SignUpForm.SetHandle(h.Account().SignUpEndpoint))
	bot.SetCommand(view.SwitchMemberCmd, h.Account().SwitchMemberEndpoint, h.CheckAuth)

	bot.SetCommand(view.CheckBalanceCmd, h.Account().GetBalanceEndpoint, h.CheckAuth)
	bot.SetCommand(view.PaymentCmd, h.Account().PaymentEndpoint, h.CheckAuth)
	return nil
}

func assets(filename string) string {
	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Dir(f)
	return filepath.Join(dir, "../../../assets/image/"+filename+".png")
}
