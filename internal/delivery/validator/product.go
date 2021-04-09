package validator

import "github.com/thedevsaddam/govalidator"

var CreateProductRules = govalidator.MapData{
	"name":     []string{"required", "max:20"},
	"price":    []string{"gtzero_decimal"},
	"cost":     []string{"gtzero_decimal"},
	"quantity": []string{"required"},
	"imageURL": []string{"required", "max:80"},
	"category": []string{"product_category"},
}

var UpdateProductRules = govalidator.MapData{
	"name":     []string{"max:20"},
	"imageURL": []string{"max:80"},
}
