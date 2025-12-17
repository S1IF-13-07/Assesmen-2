package main

import "fmt"

func main() {
	var N int
	for {
		fmt.Print("Masukkan tinggi pohon > 8: ")
		fmt.Scan(&N)
		if N > 8 {
			break
		}
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for j := 1; j < N; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}
