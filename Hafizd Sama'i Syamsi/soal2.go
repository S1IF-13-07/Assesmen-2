package main
import"fmt"
func main() {
	menu := [] string{
		"Little Trees",
		"Lap Microfiber",
		"Cover Steer",
		"Sponge Cuci Mobil",
	}
	harga := [] int{
		35000,
		25000,
		150000,
		10000,
	}
	fmt.Println("=== DAFTAR PRODUK TOKO BUDI===")
	for i := 0; i < len(menu); i++ {
		fmt.Printf("%d. %-18s - Rp%d\n", i+1, menu[i], harga[i])
	}
	var pilih, jumlah int
	fmt.Print("Pilih produk (1-4)")
	fmt.Scan(&pilih)

	fmt.Print("Masukan jumlah beli")
	fmt.Scan(&jumlah)

	total := harga[pilih-1] * jumlah
	fmt.Println("=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk : ", menu[pilih-1])
	fmt.Println("Harga : Rp", harga[pilih-1])
	fmt.Println("JUmlah :", jumlah)
	fmt.Println("Total : Rp", total)

	fmt.Println()
	fmt.Println("=== Code Excution Successful ===")
}