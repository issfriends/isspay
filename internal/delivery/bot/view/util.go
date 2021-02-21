package view

import "github.com/line/line-bot-sdk-go/linebot"

func NewPBAction(label, data string) *linebot.PostbackAction {
	return &linebot.PostbackAction{
		Label:       label,
		Data:        data,
		DisplayText: label,
	}
}

func quickRepliesView(title string, replies ...*linebot.QuickReplyButton) linebot.SendingMessage {
	items := linebot.NewQuickReplyItems(
		replies...,
	)
	msg := &linebot.TextMessage{Text: title}
	msg.WithQuickReplies(items)
	return msg
}

var DefaultMenu = linebot.RichMenu{
	Size: linebot.RichMenuSize{
		Width:  2500,
		Height: 1686,
	},
	Name:        "menu",
	ChatBarText: "menu",
	Areas: []linebot.AreaDetail{
		linebot.AreaDetail{
			Bounds: linebot.RichMenuBounds{
				X: 0, Y: 17, Width: 830, Height: 830,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypePostback,
				Data: MenuCmd.With("menu=wallet"),
			}},
		linebot.AreaDetail{
			Bounds: linebot.RichMenuBounds{
				X: 840, Y: 12, Width: 830, Height: 830,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypePostback,
				Data: MenuCmd.With("menu=shop"),
			},
		},
		linebot.AreaDetail{
			Bounds: linebot.RichMenuBounds{
				X: 1670, Y: 8, Width: 830, Height: 830,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypePostback,
				Data: MenuCmd.With("menu=function"),
			}},
	},
}
