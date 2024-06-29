package main

import "fmt"

func getCompleteName() (firstname, middleName, lastName string) {
	firstname = "muhammad"
	middleName = "daffa"
	lastName = "fauzan"

	return firstname, middleName, lastName
}

func main(){
	firstname, middleName, lastName := getCompleteName()
	fmt.Println(firstname, middleName, lastName)
}