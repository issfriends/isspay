package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/pkg/crypto"
)

type AccountDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error

	CreateAccount(ctx context.Context, account *model.Account) error
	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
	GetWallet(ctx context.Context, q *query.GetWalletQuery) error
}

type AccountServicer interface {
	Login(ctx context.Context, email, password string) (*crypto.Token, error)
	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
	SignUpByChatbot(ctx context.Context, account *model.Account) error
}

type AccountSvc struct {
	AccountDatabaser
}

func (svc AccountSvc) SignUpByChatbot(ctx context.Context, acc *model.Account) error {
	acc.UID = uuid.New().String()
	acc.Membership = value.NormalUser
	acc.Wallet = &model.Wallet{
		UID: uuid.New().String(),
	}

	if err := svc.AccountDatabaser.CreateAccount(ctx, acc); err != nil {
		return err
	}

	return nil
}

func (svc AccountSvc) Login(ctx context.Context, email, password string) (*crypto.Token, error) {

	return nil, nil
}
