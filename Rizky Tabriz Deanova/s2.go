package main

import "fmt"

func main() {
	var menu string
	var porsi int
	fmt.Print("Pilih menu: ")
	fmt.Scan(&menu)
	fmt.Print("Banyaknya porsi: ")
	fmt.Scan(&porsi)
	switch menu {
	case "LittleTrees":
		menu = "1"
	case "LapMicrofiber":
		menu = "2"
	case "CoverSteer":
		menu = "3"
	case "SpongeCuciMobil":
		menu = "4"
	default:
		fmt.Println("Menu tidak tersedia")
		return
	}
	fmt.Println("----------------------------")
	fmt.Println("Pilih produk (1-4): ", menu)
	fmt.Println("Banyaknya porsi: ", porsi)
}
