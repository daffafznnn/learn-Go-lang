package main

import "fmt"

func endApp(){
	fmt.Println("End App")
	// recover
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}
}

func main(){
	runApp(true)
	fmt.Println("Muhammad Daffa Fauzan")
}