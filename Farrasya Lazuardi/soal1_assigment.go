package main
import "fmt"

func main() {
	var X int
	for {
		fmt.Print("Masukkan tinggi pohon (X > 8): ")
		fmt.Scan(&X)
		if X > 8 {
			break
		}
		fmt.Println("Tinggi pohon harus lebih dari 8!")
	}
	for i := 1; i <= X; i++ {
		for j := 1; j <= X-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
