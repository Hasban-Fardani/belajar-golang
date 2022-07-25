package main

import "fmt"

func main() {
	// constant hampir sama dengan variabel
	// constant tidak dapat diubah nilainya sedetelah deklarasi
	// constant harus langsung dideklarasi nilainya saat ditulis
	// constant tidak akan error jika tidak digunakan
	// conts a
	// a = 10 (bakal error)

	const name = "hasban"
	const age int8 = 16
	const (
		tanggal      = 26
		bulan   int8 = 7
		tahun        = 2022
	)
	// name = "Hasban Fardani" (akan error)
	fmt.Println("nama: ", name)
}
