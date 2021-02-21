package main

import (
	"log"

	"github.com/issfriends/isspay/cmd"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{}

	root.AddCommand(cmd.Server)

	if err := root.Execute(); err != nil {
		log.Panicf("main: root command exec failed, err:%+v", err)
	}
}

// const (
// 	ChAccessToken = "yiik0wtfkny2kutyf7DOJxV1A3HOg9ApYClmlYlBb6Egw/R/vMmkLBnfQn2hFZqghban9X614gvlD/V/OGB4Hz6EXBpAmBqXPGjzamxq/eE8G1/EnlK62Vy8TD+Vr74hklVG+1IrV8Um/AFoBWbAgAdB04t89/1O/w1cDnyilFU="
// 	ChSecret      = "78d9fe6e1bbaaead2da4a92409e84255"
// )

// func main() {
// 	config := &chatbot.Config{
// 		AccessToken: ChAccessToken,
// 		Secret:      ChSecret,
// 	}
// 	wh, err := chatbot.NewLineBot(config)
// 	if err != nil {
// 		log.Panicf("init bot failed, err:%+v", err)
// 	}
// 	err = wh.SetMenu(view.DefaultMenu, "./assets/image/linebot_menu.jpg")
// 	if err != nil {
// 		log.Panicf("set menu failed, err:%+v", err)
// 	}

// 	handler := &bot.Handler{}
// 	bot.Routes(wh, handler)

// 	e := echo.New()
// 	e.GET("/ping", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "pong")
// 	})
// 	e.POST("/webhook", wh.Webhook)
// 	if err := e.Start(":8080"); err != nil {
// 		log.Panicf("init server failed, err:%+v", err)
// 	}

// }
