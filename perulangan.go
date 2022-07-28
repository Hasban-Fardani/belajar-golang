package main

import "fmt"

func main() {
	fmt.Println("===== Perulangan =====")
	var a int
	var b = 0
	fmt.Print("masukkan angka: ")
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		fmt.Printf("\nperulangan ke-%d", i)
	}
	fmt.Println("\n======================")

	for a != 0 {
		fmt.Println("ini infinite loop selama a!=0")
		fmt.Println("loop ke", b)
		b += 1
		fmt.Scanln(&a)
	}
}
