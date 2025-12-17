package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		for k := 1; k <= a-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 2; i <= a/2; i++ {
		for k := 1; k <= a-1; k++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}
