
package main

import (
	"fmt"
)

func main() {
	// Data produk
	products := []string{
		"Little Trees",
		"Lap Microfiber",
		"Cover Steer",
		"Sponge Cuci Mobil",
	}

	prices := []int{
		35000,
		25000,
		150000,
		10000,
	}

	// Tampilkan daftar produk
	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	for i := 0; i < len(products); i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, products[i], prices[i])
	}

	// Input pilihan produk
	var choice int
	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&choice)

	// Input jumlah beli
	var qty int
	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&qty)

	// Ambil data produk
	index := choice - 1
	productName := products[index]
	price := prices[index]
	total := price * qty

	// Output struk
	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk :", productName)
	fmt.Println("Harga  : Rp", price)
	fmt.Println("Jumlah :", qty)
	fmt.Println("Total  : Rp", total)
}
