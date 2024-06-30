package main

import "fmt"

// HashName adalah interface dengan satu metode GetName yang mengembalikan string.
type HashName interface {
  GetName() string
}

// sayHelloI mencetak pesan "Hello" diikuti dengan nama yang diperoleh dari metode GetName dari interface HashName.
func sayHelloI(value HashName) {
	fmt.Println("Hello", value.GetName())
}

// Person adalah struct dengan satu field Name.
type Person struct {
	Name string
}

/* GetName mengimplementasikan metode GetName dari interface HashName untuk struct Person,
 mengembalikan nilai field Name dari Person. */
func (person Person) GetName() string {
	return person.Name
}

// Animal adalah struct dengan satu field Name.
type Animal struct {
	Name string
}

/** GetName mengimplementasikan metode GetName dari interface HashName untuk struct Animal,
 mengembalikan nilai field Name dari Animal. */
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// Membuat instance dari struct Person dengan nama "Andri"
	person := Person{Name: "Andri"}
	// Memanggil fungsi sayHelloI dengan parameter person
	sayHelloI(person)

	// Membuat instance dari struct Animal dengan nama "Kucing"
	animal := Animal{Name: "Kucing"}
	// Memanggil fungsi sayHelloI dengan parameter animal
	sayHelloI(animal)
}
