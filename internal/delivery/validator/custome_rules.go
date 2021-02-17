package validator

import (
	"fmt"
	"net/url"

	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
	"github.com/vx416/govalidator"
)

func init() {
	govalidator.AddCustomRule("gtzero_decimal", func(field string, rule string, message string, value interface{}) error {
		d := value.(decimal.Decimal)
		if d.IsZero() || d.IsNegative() {
			return fmt.Errorf("The %s field must be greater than zero", field)
		}
		return nil
	})

	govalidator.AddCustomRule("product_category", func(field string, rule string, message string, val interface{}) error {
		c, ok := val.(value.ProductCategory)
		if !ok {
			return fmt.Errorf("The %s field must be product category", field)
		}
		if c != value.Drink && c != value.Snake {
			return fmt.Errorf("The %s field must be snake(1) or drink(2)", field)
		}

		return nil
	})
}

// Validate 給一個 struct 用 json tag 來 mapping rules 中的 key， message 可為 nil 會有預設的錯誤訊息
func Validate(data interface{}, rules map[string][]string, messages map[string][]string) (bool, url.Values) {
	opts := govalidator.Options{
		Data:  data,
		Rules: rules,
	}

	if messages != nil {
		opts.Messages = messages
	}

	v := govalidator.New(opts)
	v.SetTagIdentifier("json")

	e := v.ValidateStruct()

	return len(e) == 0, e
}
