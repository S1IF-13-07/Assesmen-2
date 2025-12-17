package main
import "fmt"
func main() {
	var n int

	fmt.Print("Masukkan tinggi pohon (N): ")
	fmt.Scan(&n)

	if n >= 8 {
		fmt.Println("Tinggi pohon terlalu besar, silakan kurangi!")
		return
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= (2*i - 1); k++ {
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