package restful

import (
	"errors"
	"net/http"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/delivery/restful/view"
	"github.com/issfriends/isspay/internal/delivery/validator"
	goerr "github.com/issfriends/isspay/pkg/goerr"
	"github.com/labstack/echo/v4"
	"github.com/vx416/gox/resperr"
)

type InventoryHandler struct {
	*Handler
}

// BatchCreateProductsEndpoint batch products endpoint
func (h *InventoryHandler) BatchCreateProductsEndpoint(c echo.Context) error {
	var (
		err         error
		ctx         = c.Request().Context()
		productsReq = &view.BatchCreateProductsReq{}
	)
	err = productsReq.Bind(c)
	if err != nil {
		return err
	}

	if len(productsReq.Products) > 0 {
		if err = h.svc.Inventory.BatchCreateProducts(ctx, productsReq.Products); err != nil {
			if errors.Is(err, goerr.ErrParitalUnprocessable) {
				resperr.WithDetails(err, productsReq.InvalidMsgs)
			}
			return err
		}
	}

	if len(productsReq.InvalidMsgs) != 0 {
		err = goerr.ErrParitalUnprocessable
		resperr.WithDetails(err, productsReq.InvalidMsgs)
		return err
	}

	return c.NoContent(http.StatusCreated)
}

// UpdateProductEndpoint update product endpoint
func (h *InventoryHandler) UpdateProductEndpoint(c echo.Context) error {
	var (
		err         error
		ctx         = c.Request().Context()
		updatedData = &model.Product{}
	)

	pUID := c.Param("productUID")
	if err != nil {
		return err
	}
	q := query.GetProductQuery{
		UID: pUID,
	}

	if err := c.Bind(updatedData); err != nil {
		return err
	}

	if ok, _ := validator.Validate(updatedData, validator.UpdateProductRules, nil); !ok {
		return goerr.ErrUnprocessableEntity
	}

	if !updatedData.Price.IsZero() && !updatedData.Cost.IsZero() {
		if updatedData.Price.LessThan(updatedData.Cost) {
			return goerr.ErrUnprocessableEntity
		}
	}

	if err := h.svc.Inventory.UpdateProduct(ctx, &q, updatedData); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// ListProductsEndpoint list products endpoint
func (h *InventoryHandler) ListProductsEndpoint(c echo.Context) error {
	var (
		err      error
		ctx      = c.Request().Context()
		dataList = &view.DataList{}
		q        = query.ListProductsQuery{}
	)

	if err := c.Bind(q); err != nil {
		return err
	}

	total, err := h.svc.Inventory.ListProducts(ctx, &q)
	if err != nil {
		return err
	}

	dataList.Data = q.Data
	dataList.SetTotals(q.Page, q.PerPage, total)
	return c.JSON(http.StatusOK, dataList)
}

func (h *InventoryHandler) DeleteProductEndpoint(c echo.Context) error {
	var (
		err error
		ctx = c.Request().Context()
		q   = query.GetProductQuery{}
	)

	q.UID = c.Param("productUID")
	if q.UID == "" {
		return goerr.ErrUnprocessableEntity
	}

	if err = h.svc.Inventory.DeleteProduct(ctx, &q); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
