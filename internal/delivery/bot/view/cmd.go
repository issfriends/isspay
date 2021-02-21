package view

import (
	"github.com/issfriends/isspay/pkg/chatbot"
)

const (
	MenuCmd chatbot.Command = "menu"

	SignUpCmd       chatbot.Command = "signUp"
	MeCmd           chatbot.Command = "me"
	SwitchMemberCmd chatbot.Command = "switchMember"
	BroadcastCmd    chatbot.Command = "broadcast"

	PurchaseProductCmd chatbot.Command = "purchaseProduct"
	ListProductsCmd    chatbot.Command = "listProducts"
	CancelOrderCmd     chatbot.Command = "cancelPurchase"

	ListOrdersCmd   chatbot.Command = "listOrders"
	CheckBalanceCmd chatbot.Command = "checkBalance"
	PaymentCmd      chatbot.Command = "payment"
	TransferCmd     chatbot.Command = "transfer"

	IssIntroCmd chatbot.Command = "IssIntroCmd"
	VisitorCmd  chatbot.Command = "visitorCmd"

	NotifyOutOfStockCmd chatbot.Command = "outOfStock"
)
