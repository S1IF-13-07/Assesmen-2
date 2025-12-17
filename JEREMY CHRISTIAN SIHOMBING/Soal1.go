package main
import "fmt"
func main() {

	var asterik int
	fmt.Print("Masukkan ketinggian pohon natal lebih dari 8 asterik: ")
	fmt.Scan(&asterik)
	for i := 1; i <= asterik; i++ {
		for j := 0; j < asterik-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (2*i)-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < asterik-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}