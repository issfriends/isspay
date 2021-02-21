package view

import (
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/line/line-bot-sdk-go/linebot"
)

func GetWalletView(wallet *model.Wallet, replies ...*linebot.QuickReplyButton) linebot.SendingMessage {
	replies = append(replies, walletOptions(100)...)
	return QuickRepliesView("", replies...)
}

func walletOptions(amount int64) []*linebot.QuickReplyButton {
	return []*linebot.QuickReplyButton{
		linebot.NewQuickReplyButton("", NewPBAction("我要還款", PaymentCmd.With("amount=%d", amount))),
	}
}
