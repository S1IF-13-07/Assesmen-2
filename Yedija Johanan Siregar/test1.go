package main
import "fmt"
func main() {
  var n int
  for {
  fmt.Print("Masukkan ketinggian pohon (N > 8): ")
  fmt.Scan(&n)

    if n > 8 {
      break
  }

    fmt.Println("Tinggi pohon harus lebih dari 8!")
  }
    for i := 0; i < n-1; i++ {
    for j := 0; j < n-i-1; j++ {
    fmt.Print(" ")
  }

    for j := 0; j < 2*i+1; j++ {
    fmt.Print("*")
  }
    fmt.Println()
}
  for i := 0; i < 2; i++ {
  for j := 0; j < n-2; j++ {
  fmt.Print(" ")
}

  fmt.Println("|||")
}
  fmt.Println("Merry Christmas!!")
}
