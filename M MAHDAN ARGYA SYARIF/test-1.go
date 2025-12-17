package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		padding := strings.Repeat(" ", n-i)
		asterisks := strings.Repeat("*", 2*i-1)
		fmt.Println(padding + asterisks)
	}

	trunk := strings.Repeat(" ", n-1) + "|"
	fmt.Println(trunk)
	fmt.Println(trunk)
	fmt.Println("PADORU PADORUUUUU")
}
