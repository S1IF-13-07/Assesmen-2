package main
import "fmt"
func main() {
  var pilihan, jumlah, total int
  
  fmt.Println("1. Empal Gentong Biasa - Rp22.000")
  fmt.Println("2. Empal Asem - Rp22.000")
  fmt.Println("3. Sate Kambing Muda - Rp40.000")
  fmt.Println("4. Nasi Lengko Pagengan - Rp15.000")
  
  fmt.Print("\nPilih menu (1-4): ")
  fmt.Scan(&pilihan)
  fmt.Print("Jumlah: ")
  fmt.Scan(&jumlah)

  if pilihan == 1 || pilihan == 2 {
    total = 22000 * jumlah
  } else if pilihan == 3 {
    total = 40000 * jumlah
  } else {
    total = 15000 * jumlah
  }
fmt.Printf("\nTotal: Rp %d\n", total)
}
