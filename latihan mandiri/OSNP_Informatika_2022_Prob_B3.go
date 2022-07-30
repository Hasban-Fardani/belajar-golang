// masih salah
package main

import "fmt"

func main() {
	banyak := 0
	uang := 0
	kandungan := 0
	H := 0
	K := 0
	D := 0

	fmt.Scanln(&banyak, &uang)
	for i := 0; i < banyak; i++ {
		fmt.Scanln(&H, &K, &D)
		kandungan += K
	}
	fmt.Println(kandungan)
}
