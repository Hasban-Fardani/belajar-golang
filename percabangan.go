package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("masukkan nilai: ")
	fmt.Scanln(&nilai)

	// contoh if else
	if nilai%2 == 0 {
		fmt.Println("angka", nilai, "adalah genap")
	} else {
		fmt.Println("angka", nilai, "adalah ganjil")
	}

	// contoh if elseif dan else
	if nilai < 10 {
		fmt.Println("jenis angka: satuan")
	} else if nilai < 20 {
		fmt.Println("jenis angka: belasan")
	} else if nilai < 100 {
		fmt.Println("jenis angka: puluhan")
	} else if nilai < 1000 {
		fmt.Println("jenis angka: ratusan")
	} else {
		fmt.Println("jenis angka: lebih dari ratusan")
	}

}
