package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c float64

	fmt.Println("=== Penentu Jenis Segitiga ===")

	for {
		fmt.Print("Masukkan tiga sisi (a b c): ")
		_, err := fmt.Scan(&a, &b, &c)

		if err != nil || a <= 0 || b <= 0 || c <= 0 {
			fmt.Println("\nInputnya ngga valid wok. Pastiin semua sisi bilangan bulat positif.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		break
	}

	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("\nBukan segitiga")
		os.Exit(0)
	}

	ia, ib, ic := int(a), int(b), int(c)

	isSikuSiku := (ia*ia+ib*ib == ic*ic) ||
		(ia*ia+ic*ic == ib*ib) ||
		(ib*ib+ic*ic == ia*ia)

	if ia == ib && ib == ic {
		fmt.Println("\nSegitiga sama sisi")

	} else if isSikuSiku {
		isSamaKaki := (ia == ib) || (ia == ic) || (ib == ic)

		if isSamaKaki {
			fmt.Println("\nSegitiga siku-siku, Segitiga sama kaki")
		} else {
			fmt.Println("\nSegitiga siku-siku")
		}

	} else if ia == ib || ia == ic || ib == ic {
		fmt.Println("\nSegitiga sama kaki")

	} else {
		fmt.Println("\nSegitiga sembarang")
	}
}
