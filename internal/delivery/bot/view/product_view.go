package view

import (
	"fmt"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/line/line-bot-sdk-go/linebot"
)

var defaultImageURL = "https://i.imgur.com/Lnc1bJx.png"

func ProductsMenuMsg(products []*model.Product) linebot.SendingMessage {
	template := &linebot.CarouselTemplate{
		Columns: make([]*linebot.CarouselColumn, 0, len(products)),
	}
	for _, product := range products {
		template.Columns = append(template.Columns, productToCarouselColumn(product))
	}
	return &linebot.TemplateMessage{
		Template: template,
		AltText:  "hello",
	}
}

func productToCarouselColumn(product *model.Product) *linebot.CarouselColumn {
	productPBData := PurchaseProductCmd.With("productID=%d&quantity=%d", product.ID, 1)
	imageURL := product.ImageURL
	if imageURL == "" {
		imageURL = defaultImageURL
	}
	return &linebot.CarouselColumn{
		Title:             product.Name,
		Text:              fmt.Sprintf("數量:%d, 價格:%s", product.Quantity, product.Price.StringFixed(2)),
		ThumbnailImageURL: imageURL,
		Actions: []linebot.TemplateAction{
			&linebot.PostbackAction{
				Label: "購買",
				Data:  productPBData,
			},
		},
	}
}

func OrderCancelBtn(orderUID string) *linebot.QuickReplyButton {
	return linebot.NewQuickReplyButton("", &linebot.PostbackAction{
		Label: "取消購買", Data: CancelOrderCmd.With("orderUID=%s", orderUID),
	})
}
