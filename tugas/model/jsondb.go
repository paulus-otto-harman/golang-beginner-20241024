package model

type JsonDb interface {
	Create()
	Retrieve() interface{}
}
