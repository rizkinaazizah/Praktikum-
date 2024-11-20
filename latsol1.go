package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	motor := n/2+1
	if n%2 == 0 {
		motor = n/2
	}
		fmt.Println(motor)

	}

