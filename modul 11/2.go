package main

import "fmt"

func main() {
	var kendaraan string
	var durasi, tarif int
	fmt.Scan(&kendaraan, &durasi)

	if durasi < 1{
		durasi = 1
	}
	switch{
	case kendaraan == "motor" :
		tarif = durasi*2000
	case kendaraan == "mobil" :
		tarif = durasi*5000
	case kendaraan == "truk"  :
		tarif = durasi*8000
	}
	fmt.Printf("Rp %d", tarif)
}
