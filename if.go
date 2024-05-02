package main

import "fmt"

func main(){
	name := "asep"

	if name == "daffa"{
		fmt.Println("benar anda adalah daffa")
	} else if name == "asep"{
		fmt.Println("Hallo asep")
	}else{
		fmt.Println("salah anda bukan daffa")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}