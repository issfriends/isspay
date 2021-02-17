package chatbot

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Config struct {
	AccessToken string
	Secret      string
}

func (config *Config) InitClient() (*linebot.Client, error) {
	return linebot.New(config.Secret, config.AccessToken)
}

func NewLineBot(config *Config) (ChatBot, error) {
	client, err := config.InitClient()
	if err != nil {
		return nil, err
	}

	bot := &lineBot{
		LineBotClient:  &LineBotClient{Client: client},
		routes:         make(map[Command]MsgHandle),
		forms:          make(map[Command]*MsgForm),
		activeForms:    &Forms{},
		errHandle:      DefaultErrorHandle,
		notFoundHandle: DefaultNotFoundHandle,
	}

	return bot, nil
}

// ChatBot chat book handler
type ChatBot interface {
	Webhook(c echo.Context) error
	Use(wrappers ...HandleWrapper)
	SetCommand(cmd Command, handle MsgHandle)
	SetForm(form *MsgForm)
	SetMenu(menu linebot.RichMenu, imagePath string) error
	HandleMsg(msg *MsgContext) error
}

type (
	MsgHandle     func(c *MsgContext) error
	ErrHandle     func(err error, c *MsgContext)
	HandleWrapper func(next MsgHandle) MsgHandle
)

type lineBot struct {
	menuID string
	*LineBotClient
	routes         map[Command]MsgHandle
	wrappers       []HandleWrapper
	errHandle      ErrHandle
	activeForms    *Forms
	forms          map[Command]*MsgForm
	notFoundHandle MsgHandle
}

func (bot *lineBot) Use(wrappers ...HandleWrapper) {
	if len(wrappers) > 0 {
		bot.wrappers = append(bot.wrappers, wrappers...)
	}
}

func (bot *lineBot) SetMenu(menu linebot.RichMenu, imagePath string) error {
	resp, err := bot.CreateRichMenu(menu).Do()
	if err != nil {
		return err
	}
	menuID := resp.RichMenuID
	_, err = bot.UploadRichMenuImage(menuID, imagePath).Do()
	if err != nil {
		return err
	}
	bot.menuID = menuID
	_, err = bot.SetDefaultRichMenu(menuID).Do()
	if err != nil {
		return err
	}
	return nil
}

func (bot *lineBot) SetCommand(cmd Command, handle MsgHandle) {
	if len(bot.wrappers) > 0 {
		for _, wrapper := range bot.wrappers {
			handle = wrapper(handle)
		}
	}
	bot.routes[cmd] = handle
}

func (bot *lineBot) SetErrHandle(errHandle ErrHandle) {
	bot.errHandle = errHandle
}

func (bot *lineBot) SetNotFoundHandle(handle MsgHandle) {
	bot.notFoundHandle = handle
}

func (bot *lineBot) SetForm(form *MsgForm) {
	handle := form.handle
	if len(bot.wrappers) > 0 {
		for _, wrapper := range bot.wrappers {
			handle = wrapper(handle)
		}
	}
	form.handle = handle
	bot.forms[form.cmd] = form

}

func (bot *lineBot) Webhook(c echo.Context) error {
	events, err := bot.ParseRequest(c.Request())
	if err != nil {
		return err
	}
	ctx := c.Request().Context()

	for _, event := range events {
		log.Printf("%+v", event)
		msg, err := parseMsg(ctx, bot, event)

		if err != nil {
			bot.errHandle(err, msg)
			continue
		}

		if err := bot.HandleMsg(msg); err != nil {
			bot.errHandle(err, msg)
		}
	}
	return c.NoContent(http.StatusOK)
}

func (bot *lineBot) HandleMsg(msg *MsgContext) error {
	form, ok := bot.forms[msg.Cmd]
	if ok {
		err := form.Active(msg, bot)
		if err != nil {
			return err
		}
		return nil
	}

	form, ok = bot.activeForms.GetForm(msg.GetMessengerID())
	if ok {
		done, err := form.HandleInput(msg)
		if err != nil {
			return err
		}
		if done {
			bot.activeForms.RemoveForm(msg.GetMessengerID())
		}
		return nil
	}

	handle, ok := bot.routes[msg.Cmd]
	if ok {
		err := handle(msg)
		if err != nil {
			return err
		}
	} else {
		if err := bot.notFoundHandle(msg); err != nil {
			return err
		}
	}

	return nil
}

func DefaultErrorHandle(err error, c *MsgContext) {
	if err == nil {
		return
	}
	c.ReplyMsg(TextMsg(err.Error()))
}

func DefaultNotFoundHandle(c *MsgContext) error {
	return c.ReplyMsg(TextMsg("你好"))
}
