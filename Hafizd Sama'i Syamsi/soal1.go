package main
import"fmt"
func main() {
	var N int
	fmt.Print("Masukan ketinggian piramida")
	fmt.Scan(&N)

	 
	 for i := 1; i <= N; i++ {
		for s := 0; s < N-i; s++ {
			fmt.Print(" ")
		}
		for b := 0; b < 2*i-1; b++ {
			fmt.Print("*")
		}
		fmt.Println()
	 }
	

	for i := 0; i < 2;  i++ {
		for s:= 0; s < N-1; s++ {
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}
	      
	
