package chatbot

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

type BotClient interface {
	Reply(token string, msgs ...Message) error
	Push(msgID string, msgs ...Message) error
	GetUserProfile(msgID string) (*UserProfile, error)
}

type UserProfile struct {
	MessengerID string
	Username    string
}

type Message interface {
	Message()
}

var _ BotClient = (*LineBotClient)(nil)

type LineBotClient struct {
	*linebot.Client
}

func (c *LineBotClient) Reply(token string, msgs ...Message) error {
	lineMsgs, err := c.convertMsg(msgs...)
	if err != nil {
		return err
	}

	_, err = c.ReplyMessage(token, lineMsgs...).Do()
	if err != nil {
		return err
	}
	return nil
}

func (c *LineBotClient) Push(msgID string, msgs ...Message) error {
	lineMsgs, err := c.convertMsg(msgs...)
	if err != nil {
		return err
	}

	_, err = c.PushMessage(msgID, lineMsgs...).Do()
	if err != nil {
		return err
	}
	return nil
}

func (c *LineBotClient) convertMsg(msgs ...Message) ([]linebot.SendingMessage, error) {
	lineMsgs := make([]linebot.SendingMessage, len(msgs))
	for i := range lineMsgs {
		ok := false
		lineMsgs[i], ok = msgs[i].(linebot.SendingMessage)
		if !ok {
			return nil, fmt.Errorf("convert msg to linebot.SendingMessage failed")
		}
	}
	return lineMsgs, nil
}

func (c *LineBotClient) GetUserProfile(msgID string) (*UserProfile, error) {
	resp, err := c.Client.GetProfile(msgID).Do()
	if err != nil {
		return nil, err
	}
	profile := &UserProfile{
		MessengerID: resp.UserID,
		Username:    resp.DisplayName,
	}

	return profile, nil
}
