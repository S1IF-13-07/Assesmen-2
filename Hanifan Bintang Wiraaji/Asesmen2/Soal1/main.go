package main

import "fmt"

func main() {
	var a int
	fmt.Print("Masukan ketinggian pohon : ")
	fmt.Scan(&a)
	jarak := (a - 1)
	jarakbatang := (a - 1)
	pohon := 1
	if a < 8 {
		fmt.Print("Masukan ketinggian pohon yang benar(minimal 8) : ")
		fmt.Scan(&a)
	}
	for i := 0; i < a; i++ {
		for b := (jarak); b > 0; b-- {
			fmt.Print(" ")
		}
		for j := 1; j <= pohon; j++ {
			fmt.Print("*")

		}
		for c := (jarak); c > 0; c-- {
			fmt.Print(" ")
		}
		fmt.Println("")
		jarak--
		pohon += 2
	}
	for y := 0; y <= 1; y++ {
		for d := jarakbatang; d > 0; d-- {
			fmt.Print(" ")
		}
		fmt.Print("|")
		for f := jarakbatang; f > 0; f-- {
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
