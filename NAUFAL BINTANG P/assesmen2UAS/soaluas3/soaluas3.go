package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Urutkan sisi
	if a > b { a, b = b, a }
	if b > c { b, c = c, b }
	if a > b { a, b = b, a }

	if a+b > c {
		if a == b && b == c {
			fmt.Println("Segitiga sama sisi")
		} else if a*a+b*b == c*c {
			fmt.Println("Segitiga siku-siku")
		} else if a == b || b == c {
			fmt.Println("Segitiga sama kaki")
		} else {
			fmt.Println("Segitiga sembarang")
		}
	} else {
		fmt.Println("Bukan segitiga")
	}
}