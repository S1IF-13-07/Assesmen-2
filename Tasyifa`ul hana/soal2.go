package main
import "fmt"
func main() {
	var pilih, jumlah int

	fmt.Println("=== DAFTAR BARANG ===")
	fmt.Println("1. Little Trees - 35000")
	fmt.Println("2. Lap microfiber - 25000")
	fmt.Println("3. Cover Steer - 150000")
	fmt.Println("4. Sponge cuci mobil - 10000")

	fmt.Print("Pilih barang (1-4): ")
	fmt.Scan(&pilih)

	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&jumlah)

	if pilih == 1 {
		fmt.Println("Total harga:", jumlah*35000)
	} else if pilih == 2 {
		fmt.Println("Total harga:", jumlah*25000)
	} else if pilih == 3 {
		fmt.Println("Total harga:", jumlah*150000)
	} else if pilih == 4 {
		fmt.Println("Total harga:", jumlah*10000)
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}