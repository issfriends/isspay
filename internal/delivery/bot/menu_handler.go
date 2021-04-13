package bot

import (
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/line/line-bot-sdk-go/linebot"
)

func (handler Handler) Menu() MenuHandler {
	return MenuHandler{}
}

type MenuHandler struct {
}

func (h MenuHandler) MenuEndpoint(c *chatbot.MsgContext) error {
	var msg linebot.SendingMessage

	switch c.GetValue("menu") {
	case "switchMember":
		msg = view.SwitchMemberMenuView("switch member")
	case "shop":
		msg = view.ShopMenuView("isspay shop")
	case "wallet":
		msg = view.AccountMenuView("isspay account")
	case "others":
		msg = view.OthersMenuView("isspay 功能")
	default:
		msg = chatbot.TextMsgf("menu %s not found", c.GetValue("menu"))
	}
	return c.ReplyMsg(msg)
}
