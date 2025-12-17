package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int

	for {
		fmt.Print("Masukkan tinggi pohon (N > 8): ")
		fmt.Fscan(reader, &n)

		if n > 8 {
			break
		}
		fmt.Println("Tinggi pohon harus lebih dari 8, silakan input ulang.")
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