package main
import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N >= 8 {
		fmt.Println("Tinggi pohon harus kurang dari 8")
		return
	}

	for i := 1; i <= N; i++ {
		
		for j := 1; j <= N-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}