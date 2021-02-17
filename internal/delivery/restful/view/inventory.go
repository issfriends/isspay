package view

import (
	"fmt"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/delivery/validator"
	"github.com/labstack/echo/v4"
)

type BatchCreateProductsReq struct {
	Products []*model.Product `json:"products"`

	InvalidMsgs map[string]interface{}
}

func (req *BatchCreateProductsReq) Bind(c echo.Context) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	req.InvalidMsgs = make(map[string]interface{})

	if len(req.Products) == 0 {
		return nil
	}

	validProducts := make([]*model.Product, 0, len(req.Products))
	for i, p := range req.Products {
		isValid, values := validator.Validate(p, validator.CreateProductRules, nil)

		if p.Cost.GreaterThan(p.Price) {
			isValid = false
			values.Add("cost", "product cost is greater than price")
		}

		if isValid {
			validProducts = append(validProducts, p)
		} else {
			req.InvalidMsgs[fmt.Sprintf("product_%d", i+1)] = values
		}
	}

	req.Products = validProducts
	return nil
}
