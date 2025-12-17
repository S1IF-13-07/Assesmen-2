package main
import "fmt"
func main() {
	var a, b, c int
	fmt.Print("Masukkan panjang sisi segitiga :")
	for {
		fmt.Scan(&a, &b, &c)

		
		if a > 0 && b > 0 && c > 0 {
			break
		}
	}
	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan segitiga")
		return
	}
	switch {
	case a == b && b == c:
		fmt.Println("Segitiga sama sisi")

	case a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a:
		fmt.Println("Segitiga siku-siku")

	case a == b || a == c || b == c:
		fmt.Println("Segitiga sama kaki")

	default:
		fmt.Println("Segitiga sembarang")
	}
}
