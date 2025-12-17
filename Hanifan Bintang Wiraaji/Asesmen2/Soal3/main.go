package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a == b && b == c {
		fmt.Print("Segitiga sama sisi")
	} else if a == b || b == c || c == a {
		fmt.Print("Segitiga sama kaki")
	} else if ((a*a)+(b*b) == (c * c)) || ((a*a)+(c*c) == (b * b)) || ((c*c)+(b*b) == (a * a)) {
		fmt.Print("Segitiga siku siku")
	} else if (a+b <= c) || (a+c <= b) || (b+c <= a) {
		fmt.Print("Bukan segitiga")
	} else {
		fmt.Print("Segitiga sembarang")
	}
}
