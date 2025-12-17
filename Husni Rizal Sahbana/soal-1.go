package main

import "fmt"

func main() {
	var masukan_bilangan_n int

	for {
		fmt.Print("Masukkan tinggi pohon natal : ")
		fmt.Scan(&masukan_bilangan_n)
		if masukan_bilangan_n < 8 {
			break
		}
		fmt.Println("Tolong masukan tinggi harus kurang dari 8")
	}

	for i := 1; i <= masukan_bilangan_n; i++ {
		for s := masukan_bilangan_n - i; s > 0; s-- {
			fmt.Print(" ")
		}
		for b := 1; b <= 2*i-1; b++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
