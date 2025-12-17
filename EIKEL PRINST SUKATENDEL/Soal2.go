package main

import "fmt"

func main() {

	menu := []string{
		"Empal Gentong Biasa",
		"Empal Asem",
		"Sate Kambing Muda",
		"Nasi Lengko Pagongan",
	}

	harga := []int{
		22000,
		22000,
		40000,
		15000,
	}

	var pilihMenu int
	var jumlah int

	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	for i := 0; i < len(menu); i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, menu[i], harga[i])
	}

	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&pilihMenu)

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlah)

	total := harga[pilihMenu-1] * jumlah

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Println("Menu   :", menu[pilihMenu-1])
	fmt.Println("Harga  : Rp", harga[pilihMenu-1])
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total  : Rp", total)
}