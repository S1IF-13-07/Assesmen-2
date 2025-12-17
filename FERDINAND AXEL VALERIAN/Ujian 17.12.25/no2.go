package main

import (
	"fmt"
)

func main() {
	menu := []string{
		"Empal Gentong Biasa",
		"Empal Asem",
		"Sate Kambing Muda",
		"Nasi Lengko Pagonan",
	}
	harga := []int{
		22000,
		22000,
		40000,
		15000,
	}

	fmt.Println("=== MENU EMPAL GENTONG MAS FUAD ===")
	for i := 0; i < len(menu); i++ {
		fmt.Printf("%d. %s - Rp%d\n", i+1, menu[i], harga[i])
	}
	fmt.Println("----------------------------------")

	var pilihanMenu int
	var jumlahBeli int
	var totalHargaFinal int 

	for {
		fmt.Print("Pilih menu (1-4): ")
		_, err := fmt.Scan(&pilihanMenu)

		switch {
		case err != nil:
			fmt.Println("\nInputnya ngga valid. Pastiinn maneh masukkin angka.")
		case pilihanMenu >= 1 && pilihanMenu <= 4:
			goto LanjutJumlahBeli
		default:
			fmt.Println("\nInputnya ngga valid. Masukkan nomor menu antara 1 sampai 4.")
		}
		var discard string
		fmt.Scanln(&discard)
		continue
	}

	LanjutJumlahBeli:
	for {
		fmt.Print("Masukkan jumlah beli: ")
		_, err := fmt.Scan(&jumlahBeli)

		if err != nil || jumlahBeli <= 0 {
			fmt.Println("\nInputnya ngga valid. Jumlah beli harus angka positif.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}

	indexPilihan := pilihanMenu - 1
	hargaSatuan := harga[indexPilihan]
	totalHarga := hargaSatuan * jumlahBeli
	totalHargaFinal = totalHarga 

	fmt.Println("\n=== STRUK PEMBAYARAN ===")
	fmt.Printf("Menu    : %s\n", menu[indexPilihan]) 
	fmt.Printf("Harga   : Rp %d\n", hargaSatuan)
	fmt.Printf("Jumlah  : %d\n", jumlahBeli)
	fmt.Printf("Subtotal: Rp %d\n", totalHarga)
	fmt.Println("------------------------")
	fmt.Printf("TOTAL BAYAR: Rp %d\n", totalHargaFinal)
	fmt.Println("------------------------")
	
	fmt.Println("\n--- Code Execution Successful ---")
}