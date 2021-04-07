package service

import "context"

type AuthDatabaser interface {
	CreateAccount(ctx context.Context)
	GetAccount(ctx context.Context)
}

type AuthServicer interface {
	Login(ctx context.Context)
	SignUpByChatbot(ctx context.Context)
}
