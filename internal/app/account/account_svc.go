package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/pkg/encryptor"
)

type IdentityDatabaser interface {
	CreateAccount(ctx context.Context, account *model.Account) error
	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
	ExecuteTx(ctx context.Context, callback func(ctx context.Context) error) error
}

type AccountServicer interface {
	Login(ctx context.Context, email, password string) (*encryptor.Token, error)
	SignUpByChatbot(ctx context.Context, account *model.Account) error
}

func (svc service) SignUpByChatbot(ctx context.Context, acc *model.Account) error {
	acc.UID = uuid.New().String()
	acc.Membership = value.NormalUser
	acc.Wallet = &model.Wallet{
		UID: uuid.New().String(),
	}

	if err := svc.db.CreateAccount(ctx, acc); err != nil {
		return err
	}

	return nil
}

func (svc service) Login(ctx context.Context, email, password string) (*encryptor.Token, error) {
	


	return nil, nil
}
