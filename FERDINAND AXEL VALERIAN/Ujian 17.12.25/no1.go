package main

import (
	"fmt"
)

func main() {
	var ketinggianTotal int = 0
	var inputTambahan int

	for {
		switch {
		case ketinggianTotal == 0:
			fmt.Print("Masukkin ketinggian pohon wok (N > 8): ")
		default:
			fmt.Printf("yah kurang wok, cuma %d (masih <= 8). tambahin lagi wok: ", ketinggianTotal)
		}

		_, err := fmt.Scan(&inputTambahan)
		
		if err != nil {
			fmt.Println("\nInput ngga valid. Pastiin inputnya cuma angka.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		ketinggianTotal += inputTambahan
		
		if ketinggianTotal > 8 {
			fmt.Printf("\nKetinggian totalnya %d yh. otw bikin...\n\n", ketinggianTotal)
			break
		}
	}

	for i := 1; i <= ketinggianTotal; i++ {
		for j := 0; j < ketinggianTotal-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		
		fmt.Println()
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < ketinggianTotal-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

	fmt.Println("\n- --- Code Execution Successful ---")
}