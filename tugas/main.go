package main

import (
	"20241024/tugas/model"
)

func main() {
	//var model model.Order

	//models := model.Retrieve()
	//for _, model := range models {
	//	fmt.Println("->", model)
	//}

	order1 := model.InitOrder(100)
	order1.Create()
	//
	//simpanDB(&order1)
	//
	//order1 = model.InitOrder(200)
	//
	//simpanDB(&order1)
}
