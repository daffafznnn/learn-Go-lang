package main

import "fmt"

func main(){
	names := [...]string{
		 "agus",
		 "asep",
		 "joko",
		 "budi",
		 "anies",
		 "ganjar",
	}

	slice := names[4:6]
	slice2 := names[0:4]
	slice3 := names[3:]
  fmt.Println(slice)
  fmt.Println(slice2)
  fmt.Println(slice3)

	// slice function
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:]  //sabtu dan minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)
	
	daySlice2 := append(daySlice1, "Libur baru")
	// 	days := [...]string{
	// 	"Senin",
	// 	"Selasa",
	// 	"Rabu",
	// 	"Kamis",
	// 	"Jumat",
	// 	"Sabtu",
	// 	"Minggu",
	// 	"Libur baru",
	// }
	fmt.Println(daySlice2)

// make slice
newSlice := make([]string, 2, 5)
newSlice[0] = "daffa"
newSlice[1] = "fauzan"

newData := append(newSlice, "ganteng")
// newSlice[0] = "daffa"
// newSlice[1] = "fauzan"
// newSlice[2] = "ganteng"

fmt.Println(newSlice)
fmt.Println(newData)
fmt.Println(len(newSlice))
fmt.Println(cap(newSlice))

newData[1] = "budi"

// COPY SLICE
fromSlice := days[:]
toSlice := make([]string, len(fromSlice), cap(fromSlice))

copy(toSlice, fromSlice)

fmt.Println(fromSlice)
fmt.Println(toSlice)
}