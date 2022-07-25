package main

import "fmt"

func main() {
	// konfersi tipe data numeric
	var angka64 int64
	fmt.Print("masukkan angka: ")
	fmt.Scanln(&angka64)
	fmt.Println("tipe int64:", angka64)
	var angka32 int32 = int32(angka64)
	fmt.Println("tipe int32:", angka32)
	var angka16 int8 = int8(angka64)
	fmt.Println("tipe int16:", angka16)
	var angka8 int8 = int8(angka64)
	fmt.Println("tipe int8 :", angka8)

	// konfersi tipe data byte/uint8 ke string
	var nama string
	fmt.Print("masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Println("huruf pertama:", nama[0])
	fmt.Println("huruf pertama:", string(nama[0]), "[setelah konfersi]")

}
