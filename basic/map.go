package main

import "fmt"

func main(){
	// var person map[string]string = map[string]string{}
	// person["name"] = "daffa"
	// person["address"] = "bandung"

	person := map[string]string{
		"name": "daffa",
		"address": "bandung",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	// fmt.Println(person["salah"])
	fmt.Println(person)

	// function map
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Daffa fauzan"
	book["wrong"] = "Ups"

	delete(book, "wrong")

	fmt.Println(book)
}