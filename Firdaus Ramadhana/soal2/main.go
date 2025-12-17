package main

import "fmt"

func main() {
	var menu, jumlah, harga int
	var namaBarang string

	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	fmt.Println("1. Little Trees - 35000")
	fmt.Println("2. Lap microfiber - 25000")
	fmt.Println("3. Cover Steer - 150000")
	fmt.Println("4. Sponge cuci mobil - 10000")

	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&menu)

	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&jumlah)

	switch menu {
	case 1:
		namaBarang = "Little Trees"
		harga = 35000
	case 2:
		namaBarang = "Lap microfiber"
		harga = 25000
	case 3:
		namaBarang = "Cover Steer"
		harga = 150000
	case 4:
		namaBarang = "Sponge cuci mobil"
		harga = 10000
	default:
		fmt.Println("Menu tidak tersedia")
		return
	}

	totalHarga := harga * jumlah

	fmt.Println("=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk :", namaBarang)
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total  :", totalHarga)
}
