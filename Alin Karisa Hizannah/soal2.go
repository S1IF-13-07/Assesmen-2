package main

import "fmt"

func main() {
	var pilihan, jumlah, harga int
	var namaProduk string

	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	fmt.Println("1. Little Trees      - Rp35.000")
	fmt.Println("2. Lap Microfiber    - Rp25.000")
	fmt.Println("3. Cover Steer       - Rp150.000")
	fmt.Println("4. Sponge Cuci Mobil - Rp10.000")

	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&pilihan)
	for pilihan < 1 || pilihan > 4 {
		fmt.Print("Produk tidak tersedia! Pilih produk (1-4): ")
		fmt.Scan(&pilihan)
	}

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlah)

	switch pilihan {
	case 1:
		namaProduk = "Little Trees"
		harga = 35000
	case 2:
		namaProduk = "Lap Microfiber"
		harga = 25000
	case 3:
		namaProduk = "Cover Steer"
		harga = 150000
	case 4:
		namaProduk = "Sponge Cuci Mobil"
		harga = 10000
	}

	total := harga * jumlah

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk :", namaProduk)
	fmt.Println("Harga  : Rp", harga)
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total  : Rp", total)
}
