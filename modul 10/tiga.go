package main

import "fmt"

func main() {
 var b, faktor int 
 fmt.Print("Bilangan: ")
 fmt.Scan(&b)
 fmt.Print("faktor: ")
 faktor = 0
 for i := 1; i <= b; i++ {
	if b % i == 0 {
	faktor += 1
	fmt.Printf("%d ", i)  
 	}
 }
 fmt.Println(" ")

 prima := "false"
 if faktor == 2{
	prima = "true"
 }
 fmt.Print("Prima: ", prima)
}