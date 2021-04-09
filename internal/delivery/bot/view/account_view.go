package view

import (
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/pkg/i18n"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/shopspring/decimal"
)

func GetWalletView(wallet *model.Wallet, replies ...*linebot.QuickReplyButton) (linebot.SendingMessage, error) {
	replies = append(replies, walletOptions(wallet.Amount)...)

	msg, err := i18n.ZhTW("get_balance_reply", wallet)
	if err != nil {
		return nil, err
	}
	return QuickRepliesView(msg, replies...), nil
}

func walletOptions(amount decimal.Decimal) []*linebot.QuickReplyButton {
	if amount.IsNegative() {
		return []*linebot.QuickReplyButton{
			linebot.NewQuickReplyButton("", NewPBAction("我要還款", PaymentCmd.With("amount=%s", amount.String()))),
		}
	}
	return shopOptions()
}
