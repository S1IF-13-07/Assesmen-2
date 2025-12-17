package main

import "fmt"

func main() {

	var namaMenu [4]string
	var hargaMenu [4]int

	namaMenu[0] = "Empal Gentong Blasa"
	namaMenu[1] = "Empal Asem"
	namaMenu[2] = "Sate Kambing Muda"
	namaMenu[3] = "Nasl Lengko Pagongan"

	hargaMenu[0] = 22000
	hargaMenu[1] = 22000
	hargaMenu[2] = 40000
	hargaMenu[3] = 15000

	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	for i := 0; i < 4; i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, namaMenu[i], hargaMenu[i])
	}

	var pilihan, jumlah int

	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > 4 {
		fmt.Println("Pilihan tidak ada!")
		return
	}

	fmt.Print("Masukkan jumlah beli: ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Jumlah harus lebih dari 0!")
		return
	}

	total := hargaMenu[pilihan-1] * jumlah

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Printf("Menu  : %s\n", namaMenu[pilihan-1])
	fmt.Printf("Harga : Rp %d\n", hargaMenu[pilihan-1])
	fmt.Printf("Jumlah: %d\n", jumlah)
	fmt.Printf("Total : Rp %d\n", total)
	fmt.Println("\n=== Code Execution Successful ===")
}
