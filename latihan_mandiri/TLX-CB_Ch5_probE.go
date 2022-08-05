// link https://tlx.toki.id/courses/basic/chapters/05/problems/E
package main

import "fmt"

func main() {
	var angka int
	fmt.Scanln(&angka)
	if angka < 10 {
		fmt.Println("satuan")
	} else if angka < 100 {
		fmt.Println("puluhan")
	} else if angka < 1000 {
		fmt.Println("ratusan")
	} else if angka < 10000 {
		fmt.Println("ribuan")
	} else if angka < 100000 {
		fmt.Println("puluhribuan")

	}
}
