package main

import "fmt"

func main() {
	var angka int
	fmt.Scanln(&angka) // ambil input lalu disimpan ke var angka
	fmt.Println(angka)

	// apasih kegunaan dari '&' ?
	fmt.Println(&angka, ":", angka) // akan mengeluarkan memory address dari test
	// singkatnya, '&' berguna untuk mengambil memory address dari suatu var

	// a, err := fmt.Scanln(&angka)
	// fmt.Println(angka, a, err)
}
