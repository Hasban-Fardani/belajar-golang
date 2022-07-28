package main

import "fmt"

func main() {
	var nama [2]string
	nama[0] = "Hasban" // setiap array pasti dimulai dari index ke-0
	nama[1] = "Fardani"
	fmt.Println(nama[0], nama[1])

	// membuat array secara langsung
	var angka = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(angka)
}
