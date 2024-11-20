package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	teks := "bukan"
	if n<0 && n%2 == 0 {
		teks = "genap negatif"
	}
		fmt.Println(teks)

	}

