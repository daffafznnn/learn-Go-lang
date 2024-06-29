package main

import "fmt"

func main(){
var buah [3]string
buah[0] = "apel"
buah[1] = "semangka"
buah[2] = "pisang"

fmt.Println(buah[0])
fmt.Println(buah[1])
fmt.Println(buah[2])

// array secara langsung
var values = [...]int{
	90,
	80,
	70,
}

values[0] = 100

fmt.Println(values[0])
fmt.Println("jumlah isi array values adalah", len(values))
}