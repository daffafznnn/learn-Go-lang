package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// function parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
 }

//  function return value
func getAge(age string) string {
	hello := "Usia kamu adalah" + age
	return hello
}

func main(){
	// function
	sayHello()

	// function parameter
	sayHelloTo("daffa", "fauzan")
	sayHelloTo("asep", "saepudin")

	// function return value
	age := getAge("20")
	fmt.Println(age)
}