package main

import "fmt"

func main() {
	var N int

	fmt.Print("Masukkan tinggi pohon: ")
	fmt.Scan(&N)

	if N >= 8 {
		fmt.Println("Tinggi pohon harus kurang dari 8!")
		return
	}

	for i := 1; i <= N; i++ {
		for j := i; j < N; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
