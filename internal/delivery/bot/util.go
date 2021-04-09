package bot

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/issfriends/isspay/internal/pkg/crypto"
	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/issfriends/isspay/pkg/config"
	"github.com/issfriends/isspay/pkg/i18n"
)

func GetClaims(c *chatbot.MsgContext) (*crypto.Claims, error) {
	claims, ok := c.Ctx.Value(ClaimsKey{}).(*crypto.Claims)
	if !ok {
		return nil, nil
	}
	return claims, nil
}

func Replyi18nText(c *chatbot.MsgContext, msgID string, data interface{}) error {
	msg, err := i18n.ZhTW(msgID, data)
	if err != nil {
		return err
	}
	return c.ReplyTextf(msg)
}

func (h Handler) CheckAuth(next chatbot.MsgHandle) chatbot.MsgHandle {
	return func(c *chatbot.MsgContext) error {
		var (
			msgID  = c.GetMessengerID()
			ctx    = c.Ctx
			cliams = &crypto.Claims{}
		)

		token, err := h.Auth.GetTokenByMsgID(ctx, msgID)
		if err != nil {
			return err
		}
		if token == nil {
			cliams, err = h.Auth.RefreshChatbotToken(ctx, msgID)
			if err != nil {
				return err
			}
		} else {
			_, err = jwt.ParseWithClaims(token.AccessToken, cliams, func(token *jwt.Token) (interface{}, error) {
				secret, err := crypto.Base64Decode(config.Get().Secrets.TokenSecret)
				if err != nil {
					return nil, err
				}
				return secret, nil
			})
			if err != nil {
				return err
			}
		}

		c.Ctx = context.WithValue(c.Ctx, ClaimsKey{}, cliams)
		return next(c)
	}
}
