package service

type Servicer interface {
	AccountServicer
	InventoryServicer
	TransactionServicer
}
