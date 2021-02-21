package view

import "github.com/line/line-bot-sdk-go/linebot"

func ShopMenuView(title string, replies ...*linebot.QuickReplyButton) linebot.SendingMessage {
	replies = append(replies, shopOptions()...)
	return QuickRepliesView(title, replies...)
}

func shopOptions() []*linebot.QuickReplyButton {
	return []*linebot.QuickReplyButton{
		linebot.NewQuickReplyButton("", NewPBAction("吃零食", ListProductsCmd.With("category=snake"))),
		linebot.NewQuickReplyButton("", NewPBAction("喝飲料", ListProductsCmd.With("category=drink"))),
	}
}

func AccountMenuView(title string, replies ...*linebot.QuickReplyButton) linebot.SendingMessage {
	replies = append(replies, accountOptions()...)
	return QuickRepliesView(title, replies...)
}

func accountOptions() []*linebot.QuickReplyButton {
	return []*linebot.QuickReplyButton{
		linebot.NewQuickReplyButton("", NewPBAction("訂單記錄", ListOrdersCmd.String())),
		linebot.NewQuickReplyButton("", NewPBAction("查餘額", CheckBalanceCmd.String())),
		linebot.NewQuickReplyButton("", NewPBAction("付清欠款", PaymentCmd.String())),
		linebot.NewQuickReplyButton("", NewPBAction("轉帳", TransferCmd.String())),
	}
}

func FunctionMenuView(title string, replies ...*linebot.QuickReplyButton) linebot.SendingMessage {
	replies = append(replies, functionOptions()...)
	return QuickRepliesView(title, replies...)
}

func functionOptions() []*linebot.QuickReplyButton {
	return []*linebot.QuickReplyButton{
		linebot.NewQuickReplyButton("", NewPBAction("註冊", SignUpCmd.String())),
		linebot.NewQuickReplyButton("", NewPBAction("切換會員", MenuCmd.With("menu=switchMember"))),
		linebot.NewQuickReplyButton("", NewPBAction("我是面試生", SignUpCmd.With("interviewee=yes"))),
	}
}

func SwitchMemberMenuView(title string, replies ...*linebot.QuickReplyButton) linebot.SendingMessage {
	replies = append(replies, memberOptions()...)
	return QuickRepliesView(title, replies...)
}

func memberOptions() []*linebot.QuickReplyButton {
	return []*linebot.QuickReplyButton{
		linebot.NewQuickReplyButton("", NewPBAction("校友", SwitchMemberCmd.With("membership=校友"))),
		linebot.NewQuickReplyButton("", NewPBAction("phd", SwitchMemberCmd.With("membership=phd"))),
	}
}
