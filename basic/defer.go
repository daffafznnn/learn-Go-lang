package main

import "fmt"

func longging(){
	fmt.Println("Selelsai memanggil function")
}

func runApplication(){
	defer longging()
	fmt.Println("Run application")
}

func main(){
runApplication()
}