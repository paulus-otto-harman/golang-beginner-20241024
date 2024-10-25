package screens

import (
	"errors"
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type Login struct {
	Username string
	Password string
}

func (login *Login) Render() {
	fmt.Println("Masukkan Username dan Password")
	var username interface{}
	err := errors.New("")
	for err != nil {
		username, err = gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-25s : ", "Username [0 untuk keluar]"))))
	}
	login.Username = username.(string)
	if username == "0" {
		login.Username = "0"
		return
	}
	password, _ := gola.Input(gola.Args(gola.P("label", fmt.Sprintf("%-25s : ", "Password"))))
	login.Password = password.(string)
	return
}
