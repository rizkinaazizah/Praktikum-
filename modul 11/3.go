package main

import "fmt"

func main() {
	var bilangan, hasil int
	fmt.Scan(&bilangan)
	switch{
	case bilangan % 10 == 0:
		hasil = bilangan / 10
		fmt.Println("Kategori: Bilangan kelipatan 10") 
		fmt.Printf("Hasil pemabgian antara %d /10 = %d", bilangan, hasil)
	case bilangan % 5 == 0 && bilangan != 5:
		hasil = bilangan*bilangan
		fmt.Println("Kategori: Bilangan kelipatan 5") 
		fmt.Printf("Hasil kuadrat dari  %d ^ 2 = %d", bilangan, hasil)
	case bilangan % 2 != 0 :
		hasil = bilangan + (bilangan + 1)
		fmt.Println("Kategori: Bilangan ganjil") 
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d", bilangan, bilangan+1, hasil)
	case bilangan % 2 == 0:
		hasil = bilangan * (bilangan + 1)
		fmt.Println("Kategori: Bilangan genap") 
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d", bilangan, bilangan+1, hasil)
	
	}
}
