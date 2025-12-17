
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Cek syarat segitiga
	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan segitiga")
		return
	}

	// Segitiga sama sisi
	if a == b && b == c {
		fmt.Println("Segitiga sama sisi")
		return
	}

	// Urutkan sisi (supaya mudah cek siku-siku)
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

	// Segitiga siku-siku (Pythagoras)
	if x*x+y*y == z*z {
		fmt.Println("Segitiga siku-siku")
		return
	}

	// Segitiga sama kaki
	if a == b || a == c || b == c {
		fmt.Println("Segitiga sama kaki")
		return
	}

	// Segitiga sembarang
	fmt.Println("Segitiga sembarang")
}
