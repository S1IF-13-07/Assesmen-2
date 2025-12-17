package main
import "fmt"

func main() {
	var menu int
	var jumlah int
	var harga int
	var namaMenu string

	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	fmt.Println("1. Empal Gentong Biasa  - Rp22.000")
	fmt.Println("2. Empal Asem           - Rp22.000")
	fmt.Println("3. Sate Kambing Muda    - Rp40.000")
	fmt.Println("4. Nasi Lengko Pagongan - Rp15.000")

	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&menu)

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlah)

	if menu == 1 {
		namaMenu = "Empal Gentong Biasa"
		harga = 22000
	} else if menu == 2 {
		namaMenu = "Empal Asem"
		harga = 22000
	} else if menu == 3 {
		namaMenu = "Sate Kambing Muda"
		harga = 40000
	} else if menu == 4 {
		namaMenu = "Nasi Lengko Pagongan"
		harga = 15000
	}

	fmt.Println()
	fmt.Println("=== STRUK PEMBAYARAN ===")
	fmt.Println("Menu   :", namaMenu)
	fmt.Println("Harga  : Rp", harga)
	fmt.Println("Jumlah :", jumlah)
	fmt.Println("Total  : Rp", harga*jumlah)
	fmt.Println()
	fmt.Println("=== Code Execution Successful ===")
}
