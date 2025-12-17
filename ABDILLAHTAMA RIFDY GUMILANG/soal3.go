package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int

	fmt.Print("Masukkan tiga sisi segitiga (a b c): ")
	fmt.Scan(&a, &b, &c)

	if !isTriangle(a, b, c) {
		fmt.Println("Bukan segitiga")
		return
	}

	jenis := determineTriangleType(a, b, c)
	fmt.Println(jenis)
}

func isTriangle(a, b, c int) bool {

	return (a+b > c) && (a+c > b) && (b+c > a)
}

func determineTriangleType(a, b, c int) string {

	if a == b && b == c {
		return "Segitiga sama sisi"
	}

	if a == b || a == c || b == c {
		return "Segitiga sama kaki"
	}

	sides := []int{a, b, c}
	sort.Ints(sides)

	if sides[2]*sides[2] == sides[0]*sides[0]+sides[1]*sides[1] {
		return "Segitiga siku-siku"
	}

	return "Segitiga sembarang"
}
