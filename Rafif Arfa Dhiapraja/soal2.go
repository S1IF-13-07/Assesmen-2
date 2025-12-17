package main

import "fmt"

func main() {
	var pilihan, jumlah int
	var harga, total int
	var namaProduk string

	
	fmt.Println("--- DAFTAR PRODUK TOKO BUDI ---")
	fmt.Println("1. Little Trees       - Rp35.000")
	fmt.Println("2. Lap Microfiber     - Rp25.000")
	fmt.Println("3. Cover Steer        - Rp150.000")
	fmt.Println("4. Sponge Cuci Mobil  - Rp10.000")

	
	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&pilihan) 
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
	default:
		fmt.Println("\nMaaf, pilihan produk tidak ada!")
		return 
	}

	total = harga * jumlah
	
	fmt.Println("\n--- STRUK PEMBAYARAN ---")
	fmt.Printf("Produk : %s\n", namaProduk)
	fmt.Printf("Harga  : Rp %d\n", harga)
	fmt.Printf("Jumlah : %d\n", jumlah)
	fmt.Printf("Total  : Rp %d\n", total)

}