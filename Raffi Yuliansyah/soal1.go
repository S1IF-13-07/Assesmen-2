package main

import (
	"fmt"
)

func main() {
	var n int

	for kondisi:=false; !kondisi;{
		fmt.Print("Input N: ")
		fmt.Scan(&n)
		if n > 8 {
			kondisi = true
		}
		fmt.Println("Tinggi harus lebih dari 8!")
	}
	for i := 1; i <= n; i++{
		for j:= 0; j < n-i; j++{
			fmt.Print(" ")
		}
		for k := 0; k <(2*i)-1; k++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	for i:=0; i < 2; i++{
		for j:= 0; j < n-1;j++{
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}