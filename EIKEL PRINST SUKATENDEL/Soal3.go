package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Masukkan sisi a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan sisi b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan sisi c: ")
	fmt.Scan(&c)

	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan segitiga")
		return
	}

	if a == b && b == c {
		fmt.Println("Segitiga sama sisi")
	} else if a == b || a == c || b == c {
		fmt.Println("Segitiga sama kaki")
	} else if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
		fmt.Println("Segitiga siku-siku")
	} else {
		fmt.Println("Segitiga sembarang")
	}
}