package main

import "fmt"

func main() {
	var angka int
	fmt.Scanln(&angka) // ambil input lalu disimpan ke var angka
	fmt.Println(angka)

	// apasih kegunaan dari '&' ?
	var test = 1
	fmt.Println(&test) // akan mengeluarkan memory address dari test
	// singkatnya, '&' berguna untuk mengambil memory address dari suatu var
}
