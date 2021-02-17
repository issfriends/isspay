package restful

import "github.com/issfriends/isspay/internal/app"

func New(svc *app.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

type Handler struct {
	svc *app.Service
}

func (h *Handler) Inventory() *InventoryHandler {
	return &InventoryHandler{
		Handler: h,
	}
}
