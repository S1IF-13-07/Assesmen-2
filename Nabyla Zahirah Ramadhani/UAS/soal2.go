package main
import "fmt"
func main (){
	var menu, porsi, harga int
	var nama string

	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	fmt.Println("1. Little Trees     - Rp35.000")
	fmt.Println("2. Lap Microfiber    - Rp25.000")
	fmt.Println("3. Cover Steer        - Rp150.000")
	fmt.Println("4. Sponge Cuci Mobil  - Rp.10.000")

	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&menu)
	
	for menu < 1 || menu > 4{
		fmt.Print("Pilihan tidak valid. Pilih (1-4)")
		fmt.Scan(&menu)
	}

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&porsi)

	switch menu {
	case 1 :
		nama = "Little Tress"
		harga = 35000
	case 2 :
		nama = "Lap microfiber"
		harga = 25000
	case 3 :
		nama = "Cover Steer"
		harga = 150000
	case 4 :
		nama = "Sponge cuci mobil"
		harga = 10000
	}
	total := harga * porsi 
	fmt.Println("=== SRUK PEMBAYARAN ===")
	fmt.Println("Produk : ", nama)
	fmt.Println("Harga : ", harga)
	fmt.Println("Jumlah : ", porsi)
	fmt.Println("Total : ", total)
}
