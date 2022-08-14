// entah kenapa waktu itu jalan, sekarang mah malah error
package main

import (
	mypackage "belajar_golang/belajar_package/my_package"
	"fmt"
)

func main() {
	a := mypackage.Tambah(1, 2)
	fmt.Println(a)
}
