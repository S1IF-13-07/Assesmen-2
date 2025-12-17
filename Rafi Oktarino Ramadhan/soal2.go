package main

import "fmt"

func main() {
	var pilihan, jumlah int
	var namaMenu string
	var harga int

	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	fmt.Println("1. Empal Gentong Biasa  - Rp22000")
	fmt.Println("2. Empal Asem           - Rp22000")
	fmt.Println("3. Sate Kambing Muda    - Rp40000")
	fmt.Println("4. Nasi Lengko Pagongan - Rp15000")

	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&pilihan)

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlah)

	if pilihan == 1 {
		namaMenu = "Empal Gentong Biasa"
		harga = 22000
	} else if pilihan == 2 {
		namaMenu = "Empal Asem"
		harga = 22000
	} else if pilihan == 3 {
		namaMenu = "Sate Kambing Muda"
		harga = 40000
	} else if pilihan == 4 {
		namaMenu = "Nasi Lengko Pagongan"
		harga = 15000
	} else {
		fmt.Println("Menu tidak tersedia")
		return
	}

	total := harga * jumlah

	fmt.Println()
	fmt.Println("=== STRUK PEMBAYARAN ===")
	fmt.Println("Menu   :", namaMenu)
	fmt.Println("Harga  : Rp", harga)
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total  : Rp", total)
}
