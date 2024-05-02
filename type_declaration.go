package main

import "fmt"

func main() {
	type NoKtp string

	var ktpDaffa NoKtp = "2313142432432"

	var contoh string = "2313142432432"
	var contohktp NoKtp = NoKtp(contoh)

	fmt.Println(ktpDaffa)
	fmt.Println(contohktp)

}

