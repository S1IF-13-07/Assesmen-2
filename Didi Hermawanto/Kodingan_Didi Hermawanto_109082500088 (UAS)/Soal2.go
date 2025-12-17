package main
import "fmt"
func main() {

	var pilihan, jumlah int

	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	fmt.Println("1. Little Trees       - Rp35000")
	fmt.Println("2. Lap microfiber     - Rp25000")
	fmt.Println("3. Cover Steer        - Rp150000")
	fmt.Println("4. Sponge Cuci Mobil  - Rp10000")

	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&pilihan)

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlah)

	var nama string
	var harga int

	if pilihan == 1 {
		nama = "Little Trees"
		harga = 35000
	} else if pilihan == 2 {
		nama = "Lap microfiber"
		harga = 25000
	} else if pilihan == 3 {
		nama = "Cover Steer"
		harga = 150000
	} else if pilihan == 4 {
		nama = "Sponge Cuci Mobil"
		harga = 10000
	}

	total := harga * jumlah

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk :", nama)
	fmt.Println("Harga  : Rp", harga)
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total  : Rp", total)
}
