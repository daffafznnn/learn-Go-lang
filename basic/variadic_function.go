package main

import "fmt"

// variable argumen
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main(){
	// array
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10 ))
	fmt.Println(sumAll(10, 10, 10,10, 10, 10, 10, 10, 10))

	// slice
	// fmt.Println(sumAll([]int{10, 10, 10}))
	// fmt.Println(sumAll([]int{10, 10, 10, 10, 10, 10 }))
	// fmt.Println(sumAll([]int{10, 10, 10,10, 10, 10, 10, 10, 10}))

	// slice parameter
	numbers := []int{10, 10, 10,10, 10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}