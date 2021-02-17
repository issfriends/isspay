package chatbot

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func TextMsg(text ...string) linebot.SendingMessage {
	return &linebot.TextMessage{
		Text: strings.Join(text, ","),
	}
}

func TextMsgf(f string, args ...interface{}) linebot.SendingMessage {
	return &linebot.TextMessage{
		Text: fmt.Sprintf(f, args...),
	}
}
