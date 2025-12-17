package main

import "fmt"

func main() {
	var menu, jumlah, hargaPilih int
	var menuBeli string

	fmt.Println("=== Menu Empal Gentong Mas Fuad ===")
	fmt.Println("1. Empal Gentong Biasa   - Rp 22.000")
	fmt.Println("2. Empal Asem            - Rp 22.000")
	fmt.Println("3. Sate Kambing Muda     - Rp 40.000")
	fmt.Println("4. Nasi Lengko Panggonan - Rp 15.000")
	fmt.Print("Pilih Menu (1-4): ")
	fmt.Scan(&menu)
	fmt.Print("Masukan jumlah beli: ")
	fmt.Scan(&jumlah)

	switch menu {
	case 1:
		menuBeli = "Empal Gentong Biasa"
		hargaPilih = 22000
	case 2:
		menuBeli = "Empal Asem"
		hargaPilih = 22000
	case 3:
		menuBeli = "Sate Kambing Muda"
		hargaPilih = 40000
	case 4:
		menuBeli = "Nasi Lengko Panggonan"
		hargaPilih = 15000
	default:
		fmt.Println("Masukan imput yang benar")
	}

	total := jumlah * hargaPilih
	fmt.Println("=== Struk Pembayaran ===")
	fmt.Printf("Menu   : %v\n", menuBeli)
	fmt.Printf("Harga  : Rp %v\n", hargaPilih)
	fmt.Printf("Jumlah : %v\n", jumlah)
	fmt.Printf("Total  : Rp %v\n", total)

}
