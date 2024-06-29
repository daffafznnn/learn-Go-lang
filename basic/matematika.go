package main

import "fmt"

func main(){
	var a = 10
	var b = 12
	var c = a + b
	fmt.Println(c)

	//Augmented assignments
	var i = 20

	i += 20
	i *= 2
	fmt.Println(i)

	// unary operation
	var j = 1
	j++
	j++
	j--
	fmt.Println(j)

	var g = 28
	execute := g * 2 % 5
	fmt.Println("execute g:", execute)
}