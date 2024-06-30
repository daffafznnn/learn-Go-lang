package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main(){
	address := Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}