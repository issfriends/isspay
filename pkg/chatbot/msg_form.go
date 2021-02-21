package chatbot

import (
	"sync"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Forms struct {
	sync.Map
}

func (f *Forms) PushForm(messengerID string, form *MsgForm) {
	f.Store(messengerID, form)
}

func (f *Forms) GetForm(messengerID string) (*MsgForm, bool) {
	val, ok := f.Load(messengerID)
	if !ok {
		return nil, false
	}

	return val.(*MsgForm), true
}

func (f *Forms) RemoveForm(messengerID string) {
	f.Delete(messengerID)
}

func NewForm(cmd Command, text string, expiredIn time.Duration) *MsgForm {
	return &MsgForm{
		cmd: cmd,
		text: &linebot.TextMessage{
			Text: text,
		},
		fields:    make([]string, 0, 5),
		fieldMsg:  make(map[string]linebot.SendingMessage),
		expiredIn: expiredIn,
	}
}

type MsgForm struct {
	cmd       Command
	handle    MsgHandle
	index     int32
	text      linebot.SendingMessage
	fieldMsg  map[string]linebot.SendingMessage
	fields    []string
	values    []string
	expiredIn time.Duration
	expiredAt time.Time
	mu        sync.Mutex
}

func (form *MsgForm) active() *MsgForm {
	return &MsgForm{
		cmd:       form.cmd,
		handle:    form.handle,
		text:      form.text,
		fields:    form.fields,
		values:    make([]string, len(form.fields)),
		fieldMsg:  form.fieldMsg,
		expiredIn: form.expiredIn,
		expiredAt: time.Now().Add(form.expiredIn),
	}
}

func (form *MsgForm) SetInput(field string, msg linebot.SendingMessage) *MsgForm {
	form.fieldMsg[field] = msg
	form.fields = append(form.fields, field)
	return form
}

func (form *MsgForm) SetHandle(handle MsgHandle) *MsgForm {
	form.handle = handle
	return form
}

func (form *MsgForm) Active(c *MsgContext, bot *lineBot) error {
	activedForm := form.active()

	if err := c.ReplyMsg(activedForm.text, form.fieldMsg[form.fields[0]]); err != nil {
		return err
	}

	bot.activeForms.PushForm(c.GetMessengerID(), activedForm)
	return nil
}

func (form *MsgForm) HandleInput(c *MsgContext) (bool, error) {
	form.mu.Lock()
	defer form.mu.Unlock()

	now := time.Now()
	if now.After(form.expiredAt) {
		err := c.ReplyMsg(TextMsg("timeout"))
		if err != nil {
			return false, err
		}
		return true, nil
	}

	value := c.Text()
	form.values[form.index] = value
	form.index++

	if int(form.index) >= len(form.fields) {
		formData := make(map[string]string)
		for i, field := range form.fields {
			formData[field] = form.values[i]
		}
		c.load(formData)
		return true, form.handle(c)
	}

	return false, form.pushMsgf(c, form.fields[form.index])
}

func (form *MsgForm) pushMsgf(c *MsgContext, field string) error {
	return c.ReplyMsg(form.fieldMsg[field])
}
