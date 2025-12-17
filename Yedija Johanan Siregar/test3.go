package main
import "fmt"
func main() {
  var a, b, c int
  var hasil string

  fmt.Print("Masukkan sisi a: ")
  fmt.Scan(&a)
  fmt.Print("Masukkan sisi b: ")
  fmt.Scan(&b)
  fmt.Print("Masukkan sisi c: ")
  fmt.Scan(&c)

  if a == b && b == c {
    hasil = "Segitiga sama sisi"
  } else if a == b || b == c || a == c {
    hasil = "Segitiga sama kaki"
  } else if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
    hasil = "Segitiga siku-siku"
  } else if a+b > c && a+c > b && b+c > a {
    hasil = "Segitiga sembarang"
  } else {
    hasil = "Bukan segitiga"
  }
    fmt.Println("\nHasil:", hasil)
}
