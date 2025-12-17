package main
import "fmt"

func main() {
	var menu int
	var jumlahBeli int
	var total int
	var namaMenu string
	var hargaMenu int

	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	fmt.Println("1. Empal Gentong Biasa - Rp22.000")
	fmt.Println("2. Empal Asem          - Rp22.000")
	fmt.Println("3. Sate Kambing Muda   - Rp40.000")
	fmt.Println("4. Nasi Lengko Pagongan - Rp15.000")

	fmt.Print("Pilih menu(1-4): ")
	fmt.Scan(&menu)
	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlahBeli)

	switch menu {
	case 1: 
		namaMenu = "Empal Gentong Biasa"
		hargaMenu = 22000
	case 2:
		namaMenu = "Empal Asem"
		hargaMenu = 22000
	case 3:
		namaMenu = "Sate Kambing Muda"
		hargaMenu = 40000
	case 4:
		namaMenu = "Nasi Lengko Pagongan"
		hargaMenu = 15000
	default:
		fmt.Print("tidak ada di dalam menu woii")

	}
	total = hargaMenu * jumlahBeli

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Println("Menu   : ", namaMenu)
	fmt.Println("Harga  : Rp ", hargaMenu)
	fmt.Println("Jumlah : ", jumlahBeli)
	fmt.Println("Total  : Rp ", total)
	fmt.Println("\n=== Code Execution Successful ===")
}