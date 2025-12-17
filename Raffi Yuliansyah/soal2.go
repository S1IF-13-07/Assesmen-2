package main

import "fmt"

func main(){
	var harga_menu_1, harga_menu_2, harga_menu_3, harga_menu_4 int
	var banyak_pesanan, menu_pilihan, harga_makanan, total_bayar  int
	var nama_makanan string
	harga_menu_1 = 22000
	harga_menu_2 = 22000
	harga_menu_3 = 40000
	harga_menu_4 = 15000

	fmt.Println("=== Menu Empal Gentong Mas Fuad ===")
	fmt.Println("1. Empal Gentong Biasa \t -Rp22.000")
	fmt.Println("2. Empal Asem \t\t -Rp22.000")
	fmt.Println("3. Sate Kambing Muda \t -Rp40.000")
	fmt.Println("4. Nasi Lengko Pagongan  -Rp15.000")

	fmt.Print("Pilih menu : ")
	fmt.Scan(&menu_pilihan)
	fmt.Print("Masukkan jumlah beli : ")
	fmt.Scan(&banyak_pesanan)
	
	switch(menu_pilihan){
		case 1:
			harga_makanan = harga_menu_1
			nama_makanan = "Empal Gentong Biasa"
		case 2:
			harga_makanan = harga_menu_2
			nama_makanan = "Empal Asem"
		case 3:
			harga_makanan = harga_menu_3
			nama_makanan = "Sate Kambing Muda"
		case 4:
			harga_makanan = harga_menu_4
			nama_makanan = "Nasi Lengko Pagongan"
	}

	total_bayar = harga_makanan * banyak_pesanan

	fmt.Println("=== Struk Pembayaran ===")
	fmt.Println("Menu \t: ", nama_makanan)
	fmt.Println("Harga \t: Rp", harga_makanan)
	fmt.Println("Jumlah \t: ", banyak_pesanan)
	fmt.Println("Total \t: Rp", total_bayar)




}