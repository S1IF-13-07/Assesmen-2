package main
import "fmt"
func main() {

	var menu, porsi, harga int
	var menumakanan string
	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	fmt.Println("1. Empal Gentong Biasa  - Rp22.000")
	fmt.Println("2. Empal Asem           - Rp22.000")
	fmt.Println("3. Sate Kambing Muda    - Rp40.000")
	fmt.Println("4. Nasi Lengko Pagongan - Rp15.000")
	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&menu)
	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&porsi)

	switch menu {
	case 1:
		menumakanan = "Empal Gentong Biasa"
		harga = 22000
	case 2:
		menumakanan = "Empal Asem"
		harga = 22000
	case 3:
		menumakanan = "Sate Kambing Muda"
		harga = 40000
	case 4:
		menumakanan = "Nasi Lengko Pagongan"
		harga = 15000
	default:
		harga = 0
	}
	fmt.Println("")
	fmt.Println("=== STRUK PEMBAYARAN ===")
	fmt.Println("Menu         :", menumakanan)
	fmt.Println("Harga        : Rp", harga)
	fmt.Println("Jumlah       :", porsi)
	fmt.Println("Total        : Rp", harga*porsi)
}