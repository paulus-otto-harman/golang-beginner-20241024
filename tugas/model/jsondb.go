package model

type JsonDb interface {
	GetId() int
	Save()
	Create()
	Retrieve() interface{}
}
