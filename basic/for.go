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

	  fmt.Println("Challenge")
    
    // Pola biasa (no reverse)
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Print("*")
        }
        fmt.Println() // Pindah ke baris berikutnya
    }

		// pola tangga
		 fmt.Println("Challenge")
		for i := 0; i < 6; i++ {
			for j := 0; j < i; j++ {
				fmt.Print("*")
		 }
		 fmt.Println()
		}

		  // Pola tangga reverse horizontal
    for i := 0; i < 5; i++ {
        // Loop untuk menambahkan spasi di awal baris
        for j := 5; j > i+1; j-- {
            fmt.Print(" ")
        }
        // Loop untuk mencetak bintang
        for k := 0; k <= i; k++ {
            fmt.Print("*")
        }
        // Pindah ke baris berikutnya
        fmt.Println()
    }

    // Pola terbalik secara vertikal
    fmt.Println("Challenge")

    // Loop untuk setiap baris (terbalik)
    for i := 5; i > 0; i-- {
        // Loop untuk setiap kolom di baris tersebut
        for j := 0; j < i; j++ {
            fmt.Print("*")
        }
        // Pindah ke baris berikutnya
        fmt.Println()
    }

    // Pola terbalik secara horizontal
    fmt.Println("Challenge")

    // Loop untuk setiap baris
    for i := 0; i < 5; i++ {
        // Loop untuk menambahkan spasi di awal baris
        for j := 0; j < i; j++ {
            fmt.Print(" ")
        }
        // Loop untuk mencetak bintang
        for k := 5; k > i; k-- {
            fmt.Print("*")
        }
        // Pindah ke baris berikutnya
        fmt.Println()
    }
}
