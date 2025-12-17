package main

import "fmt"

func main() {
	var n int

	
	for {
		fmt.Print("Masukkan tinggi pohon (harus lebih dari 8): ")
		fmt.Scan(&n)

		if n > 8 {
			break
		}

		fmt.Println("Tinggi pohon harus lebih dari 8. Silakan coba lagi.")
	}

	
	for i := 1; i <= n; i++ {
		
		for j := i; j < n; j++ {
			fmt.Print(" ")
		}

		
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
