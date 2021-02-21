package chatbot

import (
	"context"
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	ReplyTokenKey   = "replyToken"
	MessagenerIDKey = "messengerID"
	TextKey         = "text"
)

type Command string

func (cmd Command) String() string {
	return "command=" + string(cmd)
}

func (cmd Command) With(f string, args ...interface{}) string {
	return "command=" + string(cmd) + "&" + fmt.Sprintf(f, args...)
}

const (
	TextCmd   Command = "text"
	StickCmd  Command = "stick"
	FollowCmd Command = "follow"
)

func parseMsg(ctx context.Context, bot *lineBot, event *linebot.Event) (*MsgContext, error) {
	msgCtx := &MsgContext{
		EventType: event.Type,
		Ctx:       ctx,
		client:    bot.LineBotClient,
		FormData:  &FormData{data: make(map[string]string)},
	}
	msgCtx.store(MessagenerIDKey, event.Source.UserID)
	msgCtx.store(ReplyTokenKey, event.ReplyToken)

	switch event.Type {
	case linebot.EventTypeMessage:
		_, ok := event.Message.(*linebot.StickerMessage)
		if ok {
			msgCtx.Cmd = StickCmd
			break
		}
		msg := event.Message.(*linebot.TextMessage)
		msgCtx.Cmd = TextCmd
		msgCtx.store(TextKey, msg.Text)
	case linebot.EventTypePostback:
		dataMap, err := parsePostBackData(event.Postback.Data)
		if err != nil {

		}
		msgCtx.Cmd = Command(dataMap["command"])
		msgCtx.load(dataMap)
	case linebot.EventTypeFollow:
		msgCtx.Cmd = FollowCmd
	default:

	}
	return msgCtx, nil
}

func parsePostBackData(data string) (map[string]string, error) {
	dataPairs := strings.Split(data, "&")
	dataMap := make(map[string]string)

	for _, pair := range dataPairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {

		}
		dataMap[kv[0]] = kv[1]
	}

	return dataMap, nil
}

type BotClienter interface {
	ReplyMessage(string, linebot.Message) error
}

type MsgContext struct {
	Cmd       Command
	EventType linebot.EventType
	*FormData
	Ctx    context.Context
	client BotClient
}

func (c *MsgContext) ReplyMsg(msgs ...Message) error {
	return c.client.Reply(c.GetReplyToken(), msgs...)
}

func (c *MsgContext) PushMsgs(msgID string, msgs ...Message) error {
	return c.client.Push(msgID, msgs...)
}

func (c *MsgContext) PushTextf(f string, args ...interface{}) error {
	return c.client.Reply(c.GetReplyToken(), TextMsgf(f, args...))
}

func (c *MsgContext) GetCurrentUser() (*UserProfile, error) {
	mID := c.GetMessengerID()
	return c.client.GetUserProfile(mID)
}
