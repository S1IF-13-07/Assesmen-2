package main

import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan tinggi pohon : ")
	fmt.Scan(&x)
	for x >= 8 {
		fmt.Println("Masukan tinggi pohon (1-7) : ")
		fmt.Scan(&x)
	}

	for i := 1; i <= x; i++ {

		for s := 0; s < x-i; s++ {
			fmt.Print(" ")
		}

		for b := 0; b < 2*i-1; b++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for t := 0; t < 2; t++ {
		for s := 0; s < x-1; s++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}