package main

import "fmt"

func main(){
	name := "Asep"

	switch name {
	case "Daffa":
		fmt.Println("Hallo Daffa")
	case "Asep":
		fmt.Println("Hallo Asep")
	default: 
		fmt.Println("Lah, lu siapa?")
	}

	// switch short statement
	// switch length := len(name); length > 5{
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
  // case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// switch tanpa kondisi
		length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}