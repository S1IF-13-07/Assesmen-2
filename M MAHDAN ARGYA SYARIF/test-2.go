package main

import (
	"fmt"
	"strings"
)

func main() {
	menus := []string{
		"Empal Gentong Biasa",
		"Empal Asem",
		"Sate Kambing Muda",
		"Nasi Lengko Pagongan",
	}
	prices := []int{22000, 22000, 40000, 15000}

	fmt.Println("=== MENU POS ===")
	for i, menu := range menus {
		fmt.Printf("%d. %s - Rp %d\n", i+1, menu, prices[i])
	}

	fmt.Println(strings.Repeat("-", 40))
	var menuChoice, xontol int
	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&menuChoice)
	fmt.Print("Berapa porsi: ")
	fmt.Scan(&xontol)
	fmt.Println(strings.Repeat("-", 40))
	if menuChoice < 1 || menuChoice > 4 || xontol < 1 {
		fmt.Println("Go Fuck Yourself dude...")
		return
	}

	selectedMenu := menus[menuChoice-1]
	selectedPrice := prices[menuChoice-1]
	total := selectedPrice * xontol
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Menu        : " + selectedMenu)
	fmt.Println("Harga       : Rp " + fmt.Sprint(selectedPrice))
	fmt.Println("Porsi       : " + fmt.Sprint(xontol))
	fmt.Println("Total       : Rp " + fmt.Sprint(total))
	fmt.Println(strings.Repeat("=", 40))
}
