package view

import (
	"time"

	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/line/line-bot-sdk-go/linebot"
)

// SignUpForm sign up form message
var SignUpForm = chatbot.NewForm(SignUpCmd, "welcome iss", 60*time.Second).
	SetInput("email", chatbot.TextMsg("請輸入你的 email (如果是 iss 學生請輸入 iss email)")).
	SetInput("nickname", chatbot.TextMsg("請輸入你的 nickname")).
	SetInput("role", QuickRepliesView("請選擇你的身份", memberInputReplies()...))

func memberInputReplies() []*linebot.QuickReplyButton {
	return []*linebot.QuickReplyButton{
		linebot.NewQuickReplyButton("", linebot.NewMessageAction("master", "1")),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction("phd", "2")),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction("faculty", "3")),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction("professor", "4")),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction("alumni", "5")),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction("interviewee", "6")),
	}
}
