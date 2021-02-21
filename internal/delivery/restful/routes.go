package restful

import "github.com/labstack/echo/v4"

// Routes setup routes
func (h *Handler) Routes(serv *echo.Echo) {
	v1API := serv.Group("/api/v1")

	inventoryRoutes(h.Inventory(), v1API)
}

func inventoryRoutes(h *InventoryHandler, v1G *echo.Group) {
	v1G.POST("/products", h.BatchCreateProductsEndpoint)
	v1G.PUT("/products/:productUID", h.UpdateProductEndpoint)
	v1G.GET("/products", h.ListProductsEndpoint)
	v1G.DELETE("/products/:productUID", h.DeleteProductEndpoint)
}
