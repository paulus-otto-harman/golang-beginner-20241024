package main

import (
	"20241024/tugas/screens"
	"context"
	"fmt"
	"github.com/paulus-otto-harman/golang-module"
	"time"
)

func show(view screens.Screen) {
	gola.ClearScreen()
	view.Render()
}

func main() {
	const SessionTimeout = 5

	ctx, cancel := context.WithTimeout(context.Background(), SessionTimeout*time.Second)
	defer cancel()

	var loginScreen screens.Login
	var homeScreen screens.Home

mainLoop:
	for {
		show(&loginScreen)
		select {
		case <-ctx.Done():
			if loginScreen.Username != "0" {
				fmt.Println("Sesi telah berakhir. Silakan login kembali")
			}
			return
		default:
			switch {
			case loginScreen.Username == "0":
				break mainLoop
			case loginScreen.Username == "x" && loginScreen.Password == "x":
				show(&homeScreen)
			}
		}
	}

	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//var model model.Order

	//orders := (&model.Order{}).Retrieve().([]model.Order)
	//for _, order := range orders {
	//	fmt.Println("->", order.Id, order.Foods)
	//}

	//order1 := model.InitOrder(400)
	//fmt.Println(order1)
	//order1.Create()
	//
	//simpanDB(&order1)
	//
	//order1 = model.InitOrder(200)
	//
	//simpanDB(&order1)
}
