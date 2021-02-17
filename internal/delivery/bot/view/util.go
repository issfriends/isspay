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
		Height: 1600,
	},
	Name:        "test",
	ChatBarText: "test123",
	Areas: []linebot.AreaDetail{
		linebot.AreaDetail{
			Bounds: linebot.RichMenuBounds{
				X: 0, Y: 0, Width: 2500, Height: 800,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypePostback,
				Data: MenuCmd.With("menu=shop"),
			}},
		linebot.AreaDetail{
			Bounds: linebot.RichMenuBounds{
				X: 4, Y: 900, Width: 830, Height: 800,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypePostback,
				Data: MenuCmd.With("menu=shop"),
			},
		},
		linebot.AreaDetail{
			Bounds: linebot.RichMenuBounds{
				X: 840, Y: 900, Width: 830, Height: 800,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypePostback,
				Data: MenuCmd.With("menu=function"),
			}},
		linebot.AreaDetail{
			Bounds: linebot.RichMenuBounds{
				X: 1683, Y: 900, Width: 830, Height: 800,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypePostback,
				Data: MenuCmd.With("menu=wallet"),
			}},
	},
}
