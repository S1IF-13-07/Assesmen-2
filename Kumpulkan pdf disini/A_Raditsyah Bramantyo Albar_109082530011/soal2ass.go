package main

import "fmt"

func main() {
	var pilih int
	var jumlah int
	var total int

	fmt.Println("MENU MAS FUAD")
	fmt.Println("1. Empal Gentong Biasa - 22000")
	fmt.Println("2. Empal Asem - 22000")
	fmt.Println("3. Sate Kambing Muda - 40000")
	fmt.Println("4. Nasi Lengko Pagongan - 15000")

	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilih)

	fmt.Print("Jumlah beli: ")
	fmt.Scan(&jumlah)

	if pilih == 1 {
		total = 22000 * jumlah
		fmt.Println("Menu   : Empal Gentong Biasa")
	} else if pilih == 2 {
		total = 22000 * jumlah
		fmt.Println("Menu   : Empal Asem")
	} else if pilih == 3 {
		total = 40000 * jumlah
		fmt.Println("Menu   : Sate Kambing Muda")
	} else if pilih == 4 {
		total = 15000 * jumlah
		fmt.Println("Menu   : Nasi Lengko Pagongan")
	}

	fmt.Println("Total  : Rp", total)
}

