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
	if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
		fmt.Println("Segitiga siku-siku")
		return
	}
	if a == b || b == c || a == c {
		fmt.Println("Segitiga sama kaki")
		return
	}
	fmt.Println("Segitiga sembarang")
}