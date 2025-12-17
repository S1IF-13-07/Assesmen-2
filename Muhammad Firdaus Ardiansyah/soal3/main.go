package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Masukkan tiga bilangan: ")
	fmt.Scan(&a, &b, &c)

	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("bukan segitiga")
	} else {
		if a == b && b == c {
			fmt.Println("segitiga sama sisi")
			
		} else if a == b || a == c || b == c {
			fmt.Println("segitiga sama kaki")
			
		} else if (a*a + b*b == c*c) || (a*a + c*c == b*b) || (b*b + c*c == a*a) {
			fmt.Println("segitiga siku-siku")
			
		} else {
			fmt.Println("segitiga sembarang")
		}
	}
}