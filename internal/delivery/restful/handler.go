package restful

import "github.com/issfriends/isspay/internal/app"

func New(svc *app.App) *Handler {
	return &Handler{
		svc: svc,
	}
}

type Handler struct {
	svc *app.App
}

func (h *Handler) Inventory() *InventoryHandler {
	return &InventoryHandler{
		Handler: h,
	}
}
