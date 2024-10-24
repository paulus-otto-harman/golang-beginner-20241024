package model

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Order struct {
	Jumlah    int    `json:"jumlah"`
	CreatedAt string `json:"created_at"`
}

func InitOrder(jumlah int) Order {
	return Order{CreatedAt: time.DateTime, Jumlah: jumlah}
}

func (order *Order) Create() {
	orders := (order.Retrieve()).([]Order)
	file, err := os.Create("database/order.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	orders = append(orders, *order)

	if err := encoder.Encode(orders); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func (order *Order) Retrieve() interface{} {
	file, err := os.Open("database/order.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var orders []Order
	if err := decoder.Decode(&orders); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return orders
}
