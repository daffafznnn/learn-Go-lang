package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Printf("User %s is blacklisted\n", name)
	} else {
			fmt.Printf("User %s is registered\n", name)
	}
	
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Daffa", blacklist)
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}