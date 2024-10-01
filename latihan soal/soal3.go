package main

import (
	"fmt"
	"math"
)

func main() {
	var r, rumus float64
	fmt.Scan(&r)
	rumus = math.Pi * math.Pow(r, 2)
	fmt.Print(rumus)

}
