package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/pkg/crypto"
	"github.com/issfriends/isspay/pkg/config"
)

type AuthCacher interface {
	CacheTokenWithMsgID(ctx context.Context, token *crypto.Token, msgID string) error
}

type AuthDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error

	CreateAccount(ctx context.Context, account *model.Account) error
	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
}

type AuthServicer interface {
	// Login(ctx context.Context)
	SignUpByChatbot(ctx context.Context, account *model.Account) error
	RefreshChatbotToken(ctx context.Context, messengerID string) (*crypto.Claims, error)
}

func NewAuth(db AuthDatabaser, cache AuthCacher) AuthServicer {
	return &authSvc{
		AuthDatabaser: db,
		AuthCacher:    cache,
	}
}

type authSvc struct {
	AuthDatabaser
	AuthCacher
}

func (svc authSvc) SignUpByChatbot(ctx context.Context, account *model.Account) error {

	svc.AuthDatabaser.ExecuteTx(ctx, func(txCtx context.Context) error {
		getAccQ := &query.GetAccountQuery{
			Email: string(account.Email),
		}
		err := svc.AuthDatabaser.GetAccount(ctx, getAccQ)
		if err != nil {
			return err
		}

		if getAccQ.Data != nil {
			// duplicate error
			return errors.New("duplicated")
		}

		account.UID = uuid.New().String()

		err = svc.AuthDatabaser.CreateAccount(ctx, account)
		if err != nil {
			return err
		}

		_, err = svc.RefreshChatbotToken(ctx, account.MessengerID.String)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

func (svc authSvc) RefreshChatbotToken(ctx context.Context, messengerID string) (*crypto.Claims, error) {
	getAccQ := &query.GetAccountQuery{
		MessengerID: messengerID,
		HasWallet:   true,
	}

	err := svc.AuthDatabaser.GetAccount(ctx, getAccQ)
	if err != nil {
		return nil, err
	}

	account := getAccQ.Data

	return svc.cacheToken(ctx, account, messengerID)
}

func (svc authSvc) cacheToken(ctx context.Context, account *model.Account, msgID string) (*crypto.Claims, error) {
	claims := &crypto.Claims{
		AccountID:  account.ID,
		WalletID:   account.Wallet.ID,
		Role:       int64(account.Role),
		Membership: int64(account.Membership),
	}
	claims.ExpiresAt = time.Now().Add(30 * 24 * time.Hour).Unix()

	token, err := claims.SignToToken(config.Get().Secrets.TokenSecret)
	if err != nil {
		return nil, err
	}

	err = svc.CacheTokenWithMsgID(ctx, token, msgID)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
