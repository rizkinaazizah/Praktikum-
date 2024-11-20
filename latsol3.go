package main

import "fmt"

func main(){
	var x, y int
	var faktorX, faktorY bool
	fmt.Scan(&x, &y)

	if y%x == 0 {
		faktorX = true
	}
	if x%y == 0 {
		faktorY = true
	}
	
	fmt.Println(faktorX)
	fmt.Print(faktorY)
}