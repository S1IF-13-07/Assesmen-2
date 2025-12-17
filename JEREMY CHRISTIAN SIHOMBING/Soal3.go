package main 
import "fmt"
func main() {

	var a, b, c int
	fmt.Print("Masukkan tiga bilangan bulat positif a b dan c: ")
	fmt.Scan(&a, &b, &c)
	switch {
	case a + b <= c:
		fmt.Println("Bukan segitiga")
	case a == c:
		fmt.Println("Segitiga sama sisi")
	case (a*a)+(b*b) == (c*c):
		fmt.Println("Segitiga siku-siku")
	case a == b || b == c:
		fmt.Println("Segitiga sama kaki")
	default:
		fmt.Println("Segitiga sembarang")
	}
}