
package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan tinggi pohon (N): ")
	fmt.Scan(&n)

	// Validasi tinggi pohon
	if n >= 8 {
		fmt.Println("Tinggi pohon terlalu besar, silakan kurangi (kurang dari 8).")
		return
	}

	// Mencetak daun pohon (segitiga)
	for i := 1; i <= n; i++ {
		// Spasi kiri
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		// Asterisk
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Mencetak batang pohon (2 baris)
	for i := 0; i < 2; i++ {
		for j := 0; j < n-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}
