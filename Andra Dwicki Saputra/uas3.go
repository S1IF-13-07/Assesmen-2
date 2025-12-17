package main
import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan segitiga")
		return
	}

	if a == b && b == c {
		fmt.Println("Segitiga sama sisi")
		return
	}

	x, y, z := a, b, c
	if x > y {
		x, y = y, x
	}
	if y > z {
		y, z = z, y
	}
	if x > y {
		x, y = y, x
	}

	if x*x+y*y == z*z {
		fmt.Println("Segitiga siku-siku")
		return
	}

	if a == b || a == c || b == c {
		fmt.Println("Segitiga sama kaki")
		return
	}

	fmt.Println("Segitiga sembarang")
}