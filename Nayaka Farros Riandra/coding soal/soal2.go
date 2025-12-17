package main

import "fmt"

func main(){
	var pilihan, jumlah int

	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	fmt.Println("1. Empal Gentong Biasa  - Rp22.000")
	fmt.Println("2. Empal Asem           - Rp22.000")
	fmt.Println("3. Sate Kambing Muda    - Rp40.000")
	fmt.Println("4. Nasi Lengko Pagongan - Rp15.000")
	fmt.Print("Pilih menu cuyy (1-4): ")
	fmt.Scan(&pilihan)
	fmt.Print("beli berapa? : ")
	fmt.Scan(&jumlah)

	var namaMenu string
	var harga int

	switch pilihan {
	case 1:
		namaMenu = "Empal Gentong Biasa"
		harga = 22000
	case 2:
		namaMenu = "Empal Asem"
		harga = 22000
	case 3:
		namaMenu = "Sate Kambing Muda"
		harga = 40000
	case 4:
		namaMenu = "Nasi Lengko Pagongan"
		harga = 15000
	default:
		fmt.Println("Menu tidak valid!")
		return
	}
	
	total := harga * jumlah
	fmt.Println("=== STRUK PEMBAYARAN ===")
	fmt.Println("Menu yang dipilih: ", namaMenu)
	fmt.Println("Harga per item: ", harga)
	fmt.Println("Jumlah beli: ", jumlah)
	fmt.Println("Total pembayaran: ", total)

}