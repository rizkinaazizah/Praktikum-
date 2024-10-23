package main

import "fmt"

func main() {
	var a, b, pangkat int
	
	fmt.Scan(&a, &b)
	pangkat = 2
	 for i := 1; i <= b; i++ {
		pangkat *= a
	 }
	 fmt.Print(pangkat)

	
}