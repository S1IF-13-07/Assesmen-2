package main

import (
	"fmt"
	"strings"
)

func main() {
	var tinggi int

	fmt.Print("Masukkan tinggi pohon (kurang dari 8): ")
	fmt.Scan(&tinggi)

	for tinggi >= 8 {
		fmt.Println("Tinggi harus kurang dari 8!")
		fmt.Print("Masukkan tinggi pohon (kurang dari 8): ")
		fmt.Scan(&tinggi)
	}

	for i := 1; i <= tinggi; i++ {

		spasi := strings.Repeat(" ", tinggi-i)

		bintang := strings.Repeat("*", 2*i-1)

		fmt.Println(spasi + bintang)
	}
}
