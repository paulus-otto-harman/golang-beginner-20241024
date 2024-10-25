package screens

import (
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

type Home struct {
	Menu int
}

func (home *Home) Render() {
	fmt.Println("Menu Utama")
	fmt.Println("[1] Pemesanan")
	fmt.Println("[2] Ubah Pesanan")
	fmt.Println("[3] Menerima Pembayaran")
	fmt.Println("[4] Riwayat Pesanan")
	fmt.Println("[5] Ubah Status Pesanan")
	fmt.Println("[6] Ubah Status Pesanan")

	fmt.Println()
	fmt.Println("[0] Logout")
	menuItem, _ := gola.Input(gola.Args(gola.P("type", "number"), gola.P("label", fmt.Sprintf("%-30s : ", "Masukkan 1-6 atau 0 untuk Logout"))))
	home.Menu = menuItem.(int)
}
