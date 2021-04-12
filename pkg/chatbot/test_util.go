package chatbot

import (
	"context"

	"github.com/Pallinder/go-randomdata"
	"github.com/line/line-bot-sdk-go/linebot"
)

func TestBot() ChatBot {
	bot := &lineBot{
		routes:         make(map[Command]MsgHandle),
		forms:          make(map[Command]*MsgForm),
		activeForms:    &Forms{},
		errHandle:      DefaultErrorHandle,
		notFoundHandle: DefaultNotFoundHandle,
		LineBotClient:  &LineBotClient{},
	}

	return &mockLineBot{lineBot: bot}
}

type mockLineBot struct {
	*lineBot
}

func (mock mockLineBot) SetMenu(menu linebot.RichMenu, imagePath string) error {
	return nil
}

func TestMsgCtx(cmd Command, data map[string]string) *MsgContext {
	return &MsgContext{
		Cmd:      cmd,
		FormData: newTestFormData(data),
		client:   &MockClient{},
		Ctx:      context.Background(),
	}
}

func TestForm(cmd Command, texts ...string) []*MsgContext {
	msgs := []*MsgContext{
		{
			Cmd:      cmd,
			client:   &MockClient{},
			Ctx:      context.Background(),
			FormData: newTestFormData(make(map[string]string)),
		},
	}
	msgID := msgs[0].GetMessengerID()
	for _, text := range texts {
		formData := &FormData{data: map[string]string{TextKey: text}}
		formData.store(MessagenerIDKey, msgID)
		formData.store(ReplyTokenKey, randomdata.Alphanumeric(20))

		msgs = append(msgs, &MsgContext{
			Cmd:      TextCmd,
			FormData: formData,
			client:   &MockClient{},
			Ctx:      context.Background(),
		})
	}

	return msgs
}

func newTestFormData(data map[string]string) *FormData {
	formData := &FormData{data: data}
	formData.data[MessagenerIDKey] = randomdata.Alphanumeric(20)
	formData.data[ReplyTokenKey] = randomdata.Alphanumeric(20)
	return formData
}

func TestText(text string) *MsgContext {
	return &MsgContext{
		Cmd:      TextCmd,
		FormData: newTestFormData(make(map[string]string)),
		client:   &MockClient{},
		Ctx:      context.Background(),
	}
}

type MockClient struct {
}

func (c *MockClient) Reply(token string, msgs ...Message) error {
	return nil
}

func (c *MockClient) Push(msgID string, msgs ...Message) error {
	return nil
}

func (c *MockClient) GetUserProfile(msgID string) (*UserProfile, error) {
	profile := &UserProfile{
		MessengerID: msgID,
		Username:    randomdata.FirstName(3) + " " + randomdata.LastName(),
	}
	return profile, nil
}

func SetTestMsgID(ctx *MsgContext, msgID string) {
	ctx.testMsgID = msgID
	ctx.store(MessagenerIDKey, msgID)
}
