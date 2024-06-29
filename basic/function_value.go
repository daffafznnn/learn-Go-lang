package main

import "fmt"

func getGoodbye(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}

func main() {
	goodbye := getGoodbye
	fmt.Println(goodbye("daffa"))
}