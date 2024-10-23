package main

import "fmt"
import "math"

func main() {
	var i, r, h, n int
	var volume float64

	fmt.Scan(&n)
	for i = 1; i <= n; i += 1 {
		fmt.Scan(&r, &h)
		volume = 1.0/3.0 * math.Pi * math.Pow(float64(r), 2)* float64(h)
		fmt.Println(volume)
	}
}