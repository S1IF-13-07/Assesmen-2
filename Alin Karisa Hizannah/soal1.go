package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan tinggi pohon : ")
	fmt.Scan(&n)
	
	for n >= 8 {
		fmt.Println("Kurangi tinggi pohon!")
		fmt.Print("Masukkan tinggi pohon : ")
		fmt.Scan(&n)
	}

	for i := 1; i <= n; i++ {

		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for l := 0; l < 2; l++ {
		for j := 0; j < n-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}
