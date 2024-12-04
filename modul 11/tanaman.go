package main

import "fmt"

func main() {
	var nama_tanaman string
	fmt.Scan(&nama_tanaman)
	switch nama_tanaman{
	case "nepenthes", "drosera" :
		fmt.Println(" Termasuk tanaman karnivora")
		fmt.Println(" Asli indonesia")
	case "venus", "sarracenia" :
		fmt.Println(" Termasuk tanaman karnivora")
		fmt.Println(" Tidak asli indonesia")
	default: 
	fmt.Println(" Tidak termasuk tanaman karnivora")
	}
}