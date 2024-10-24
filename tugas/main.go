package main

import (
	"20241024/tugas/model"
	"fmt"
)

func main() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//var model model.Order

	orders := (&model.Order{}).Retrieve().([]model.Order)
	for _, order := range orders {
		fmt.Println("->", order.Id, order.Foods)
	}

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
