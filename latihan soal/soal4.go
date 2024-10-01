package main

import "fmt"

func main() {
	var F int
	fmt.Scan(&F)
	F = (F - 32) * 5 / 9
	fmt.Print(F)

}
