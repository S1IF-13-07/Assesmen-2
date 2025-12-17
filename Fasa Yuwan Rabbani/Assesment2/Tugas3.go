package main
import "fmt"

func main() {
	var a, b, c int
	var d string
	fmt.Print("Masukkan Tinggi segitiga : ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang alas segitiga : ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan panjang sisi miring segitiga : ")
	fmt.Scanln(&c)
	for {

		if a > 0 && b > 0 && c > 0 {
			if a*a+b*b == c*c {
				d = "Segitiga siku-siku"
			} else if a == b && b == c {
				d = "Segitiga sama sisi"
			} else if a == b || b == c || a == c {
				d = "Segitiga sama kaki"
			} else {
				d = "Segitiga sembarang"
			}
		} else {
			d = "Input tidak valid"
		}
		fmt.Println(d)
		break
	}
}