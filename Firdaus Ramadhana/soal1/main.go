package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan tinggi pohon: ")
	fmt.Scan(&n)

	if n >= 8 {
		fmt.Println("Tinggi pohon harus kurang dari 8")
		return
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for l := 0; l < 2; l++ {
		for m := 1; m <= n-1; m++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}
