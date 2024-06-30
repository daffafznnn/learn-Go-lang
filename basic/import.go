// /learn-golang/basic/import.go
package main

import (
    "learn-golang/basic/helper"
    "fmt"
)

func main() {
    result := helper.SayHello("Daffa")
    fmt.Println(result)
		// fmt.Println(helper.version) // tidak bisa di akses
		fmt.Println(helper.Application)
		// fmt.Println(helper.sayGoodBye("Daffa")) // tidak bisa di akses
}
