package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/pkg/crypto"
	"github.com/issfriends/isspay/pkg/config"
)

type AuthCacher interface {
	CacheToken(ctx context.Context, token *crypto.Token, fromMsgID bool) error
}

type AuthDatabaser interface {
	CreateAccount(ctx context.Context, account *model.Account) error
	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
}

type AuthServicer interface {
	// Login(ctx context.Context)
	SignUpByChatbot(ctx context.Context, email value.Email, userName, messengerID string) error
	RefreshChatbotToken(ctx context.Context, messengerID string) (*crypto.Claims, error)
}

func NewAuth(db AuthDatabaser) AuthServicer {
	return &authSvc{AuthDatabaser: db}
}

type authSvc struct {
	AuthDatabaser
	AuthCacher
}

func (svc authSvc) SignUpByChatbot(ctx context.Context, email value.Email, userName, messengerID string) error {
	getAccQ := &query.GetAccountQuery{
		Email: string(email),
	}
	err := svc.AuthDatabaser.GetAccount(ctx, getAccQ)
	if err != nil {
		return err
	}

	if getAccQ.Data != nil {
		// duplicate error
		return errors.New("duplicated")
	}

	account := &model.Account{
		Email:      string(email),
		UserName:   userName,
		Membership: value.NormalUser,
		UID:        uuid.New().String(),
		Wallet: &model.Wallet{
			UID: uuid.New().String(),
		},
	}

	err = svc.AuthDatabaser.CreateAccount(ctx, account)
	if err != nil {
		return err
	}

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

	err = svc.CacheToken(ctx, token, true)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
