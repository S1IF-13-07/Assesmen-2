package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("=== Daftar Produk Toko Budi ===")
	fmt.Println("1. Little Trees        - 35.000")
	fmt.Println("2. Lap microfiber     - 25.000")
	fmt.Println("3. Cover Steer        - 150.000")
	fmt.Println("4. Sponge cuci mobil  - 10.000")

	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&a)

	fmt.Print("Masukkan jumlah porsi: ")
	fmt.Scan(&b)

	harga := 0
	nama := ""

	switch a {
	case 1:
		harga = 35000
		nama = "Little Trees"
	case 2:
		harga = 25000
		nama = "Lap microfiber"
	case 3:
		harga = 150000
		nama = "Cover Steer"
	case 4:
		harga = 10000
		nama = "Sponge cuci mobil"
	default:
		fmt.Println("Menu tidak tersedia")
		return
	}

	total := harga * b

	fmt.Println("\n=== STRUK PEMBELIAN ===")
	fmt.Println("Produk :", nama)
	fmt.Println("Jumlah :", a)
	fmt.Println("Total  :", total)
}
