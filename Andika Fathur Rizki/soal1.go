package main

import "fmt"

func main() {
	var n int

	for {
		fmt.Print("Masukan tinggi pohon (harus lebih tinggi dari 8): ")
		fmt.Scan(&n)

		if n > 8 {
			break
		}
		fmt.Println("Tinggi pohon harus lebih dari 8. Sialahkan coba lagi. ")
	}

	for i := 1; i <= n; i++ {
		for s := 1; s <= n-i; s++ {
			fmt.Print(" ")
		}
		for b := 1; b <= 2*i-1; b++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for s := 1; s < n; s++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}
