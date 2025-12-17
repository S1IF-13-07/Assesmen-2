package main
import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	fmt.Println("1. empal gentong			- Rp. 22000")
	fmt.Println("2. empal asem				- Rp. 22000")
	fmt.Println("3. Sate Kambing			- Rp. 40000")
	fmt.Println("4. nasi lengko pagongan 	- Rp. 15000")
	fmt.Print("Pilih menu (1-4) : ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan jumlah beli : ")
	fmt.Scanln(&b)

	fmt.Println("=== STRUK PEMBAYARAN ===")

	switch a {
		case 1:
			fmt.Println("Menu	: empal gentong")
			fmt.Println("Harga	: 22000")
			fmt.Println("Jumlah	: ", b)
			fmt.Println("Total : ", 22000*b)
		case 2:
			fmt.Println("Menu	: empal asem")
			fmt.Println("Harga	: 22000")
			fmt.Println("Jumlah	: ", b)
			fmt.Println("Total  : ", 22000*b)
		case 3:
			fmt.Println("Menu	: Sate Kambing")
			fmt.Println("Harga	: 40000")
			fmt.Println("Jumlah	: ", b)
			fmt.Println("Total  : ", 40000*b)
		case 4:
			fmt.Println("Menu	: nasi lengko pagongan")
			fmt.Println("Harga	: 15000")
			fmt.Println("Jumlah	: ", b)
			fmt.Println("Total  : ", 15000*b)
	}
	
}