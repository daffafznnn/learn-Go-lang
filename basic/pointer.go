package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func main(){
	// pass by value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) //berubah menjadi Bandung

	// Pointer Operator &
	address2 := &address1

	address2.City = "Bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) //berubah menjadi Bandung
}