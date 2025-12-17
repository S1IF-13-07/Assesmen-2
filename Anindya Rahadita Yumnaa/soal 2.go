package main

import "fmt"

func main() {
	produk := []string{
		"Little Trees",
		"Lap Microfiber",
		"Cover Steer",
		"Sponge Cuci Mobil",
	}
	harga := []int{
		35000,
		25000,
		150000,
		10000,
	}
	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	for i := 0; i < len(produk); i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, produk[i], harga[i])
	}
	var pilihan int
	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&pilihan)

	var jumlah int
	fmt.Print("Masukkan jumlah pembelian: ")
	fmt.Scan(&jumlah)

	index := pilihan - 1
	total := harga[index] * jumlah

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk :", produk[index])
	fmt.Println("Harga :", harga[index])
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total :", total)

	fmt.Println("\n=== Code Execution Successful ===")
}