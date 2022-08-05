// https://tlx.toki.id/courses/basic/chapters/06/problems/E
package main

import "fmt"

func main() {
	var angka int
	pangkat := true

	fmt.Scanln(&angka)
	for angka > 1 {
		if angka%2 != 0 {
			pangkat = false
			break
		} else {
			angka /= 2
		}
	}
	if pangkat {
		fmt.Println("ya")
	} else {
		fmt.Println("bukan")
	}
}
