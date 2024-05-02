package main

import "fmt"

func main() {
	//  counter := 1

	//  for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	//  }

	//  fmt.Println("sudah selsesai")

	//  for statement

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("sudah selesai")

	//  for range

	names := []string{"daffa", "fauzan", "ganteng"}

	// manual
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	//  fmt.Println("sudah selesai")

	// automatic
	for _, name := range names {
		fmt.Println(name)
	}
}
