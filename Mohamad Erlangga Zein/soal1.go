package main
import "fmt"

func main() {
	var n int

	fmt.Print("Input ketinggian piramida (N): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (2*i)-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < n-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}