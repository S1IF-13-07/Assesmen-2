package main
import "fmt"
func main() {
	var tp int
	fmt.Scan(&tp)
	for tp >= 8 {
		fmt.Println("Terlalu tinggi, masukkan lagi")
		fmt.Scan(&tp)
	}
	for i := 1; i <= tp; i++ {
		for s := 0; s < tp-i; s++ {
			fmt.Print(" ")
		}
		for b := 0; b < 2*i-1; b++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for t := 0; t < 2; t++ {
		for s := 0; s < tp-1; s++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}