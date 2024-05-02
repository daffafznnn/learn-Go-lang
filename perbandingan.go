package main

import "fmt"

func main(){
	// string
var name1 = "daffa"
var name2 = "daffa fauzan"
var result bool = name1 != name2
fmt.Println(result)

var a = 10
var b = 20
var resultValue bool = a <= b
fmt.Println(resultValue)

// operasi boolean
var nilaiAkhir = 90
var absensi = 80

var lulusNilaiAkhir = nilaiAkhir > 80
var lulusAbsensi = absensi > 80

var lulus bool = lulusNilaiAkhir && lulusAbsensi

fmt.Println("anda", lulus)
}