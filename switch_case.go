package main

import "fmt"

func pangkat(x int, n int) int {
	var hasil = 1
	if n > 0 {
		for n > 0 {
			hasil *= x
			n--
		}
	} else {
		return 1
	}
	return hasil
}

func main() {
	var (
		kode  int
		a     int
		b     int
		hasil int
	)
	fmt.Println("operator:")
	fmt.Println("1. Tambah")
	fmt.Println("2. Kurang")
	fmt.Println("3. Kali")
	fmt.Println("4. Bagi")
	fmt.Println("5. Pangkat")
	fmt.Println("6. Modulo (%)")
	fmt.Print("masukkan nomer operator: ")
	fmt.Scanln(&kode)
	fmt.Print("masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("masukkan nilai b: ")
	fmt.Scanln(&b)

	switch kode {
	case 1:
		hasil = a + b
	case 2:
		hasil = a - b
	case 3:
		hasil = a * b
	case 4:
		hasil = a / b
	case 5:
		hasil = pangkat(a, b)
	case 6:
		hasil = a % b
	default:
		fmt.Println("\ntolong masukkan angka yang benar")
	}
	fmt.Println("hasil:", hasil)
}
