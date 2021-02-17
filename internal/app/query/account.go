package query

import "github.com/issfriends/isspay/internal/app/model"

type GetAccountQuery struct {
	Email       string
	ID          int64
	MessengerID string
	Data        *model.Account
}
