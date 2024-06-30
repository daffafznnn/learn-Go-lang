package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main(){
	var daffa Customer
	// no value
	fmt.Println(daffa)

	// bervalue
	daffa.Name = "Daffa"
	daffa.Address = "Jl. Kebon Jeruk"
	daffa.Age = 18
	fmt.Println(daffa)

	asep := Customer{
		Name: "Asep",
		Address: "Jl. Kebon Kawung",
		Age: 28,
	}
	fmt.Println(asep)

	agus := Customer{"Agus", "Jl. Durian Runtuh", 20}
	fmt.Println(agus)

	agus.sayHello("budi")
}