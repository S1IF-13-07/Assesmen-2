package main

import "fmt"

func main() {
	var pilih, jumlah int
	var harga int

	fmt.Println(" MENU EMPAL GENTONG MAS FUAD ")
	fmt.Println("1. Empal Gentong Biasa  - Rp22.000")
	fmt.Println("2. Empal Asem           - Rp22.000")
	fmt.Println("3. Sate Kambing Muda    - Rp40.000")
	fmt.Println("4. Nasi Lengko Pagongan - Rp15.000")

	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&pilih)

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlah)

	if pilih == 1 {
		harga = 22000
		fmt.Println("\n STRUK PEMBAYARAN ")
		fmt.Println("Menu   : Empal Gentong Biasa")
	} else if pilih == 2 {
		harga = 22000
		fmt.Println("\n STRUK PEMBAYARAN ")
		fmt.Println("Menu   : Empal Asem")
	} else if pilih == 3 {
		harga = 40000
		fmt.Println("\n STRUK PEMBAYARAN ")
		fmt.Println("Menu   : Sate Kambing Muda")
	} else if pilih == 4 {
		harga = 15000
		fmt.Println("\n STRUK PEMBAYARAN ")
		fmt.Println("Menu   : Nasi Lengko Pagongan")
	}

	fmt.Println("Harga  : Rp", harga)
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total  : Rp", harga*jumlah)
}
