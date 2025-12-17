package main
import "fmt"
func main() {
	var x, y, harga int
	var nama string

	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	fmt.Println("1. Little Trees     - Rp35.000")
	fmt.Println("2. Lap microfiber    -Rp25.000")
	fmt.Println("3. Cover Steer       -Rp150.000")
	fmt.Println("4. Sponge cuci mobil -Rp10.000")

	fmt.Print("Pilih produk (1-4): ")
	fmt.Scan(&x)

	for x < 1 || x > 4{
		fmt.Print("Produk tidak valid. Pilih (1-4) : ")
		fmt.Scan(&x)
	}

	fmt.Print("Masukan jumlah beli : ")
	fmt.Scan(&y)

	switch x{
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
	
	total := harga * y

	fmt.Println("=== SRUK PEMBAYARAN ===")
	fmt.Println("Produk : ", nama)
	fmt.Println("Harga : ", harga)
	fmt.Println("Jumlah : ", y)
	fmt.Println("Total : ", total)
}
