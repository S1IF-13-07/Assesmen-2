package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Masukkan tiga bilangan bulat: ")
	fmt.Scan(&a, &b, &c)

	if a+b <= c || a+c <= b || b+c <= a{
		fmt.Println("Bukan Segitiga")
		return
	}
	if a == b && b == c {
		fmt.Println("Segitiga Sama Sisi")
		return
	}
	if a == b || a == c || b == c {
		fmt.Println("Segitiga Sama Kaki")
		return
	}
	fmt.Println("Segitiga Sembarang")
}