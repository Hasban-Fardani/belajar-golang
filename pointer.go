package main

import "fmt"

func tambah(angka *int, penambah int) {
	*angka += penambah
}

func main() {
	a := 1
	b := a
	fmt.Println("a,b: ", a, b)
	b += 10
	fmt.Println("a,b: ", a, b) // nilai a tidak berubah

	x := 16
	y := &x
	fmt.Println("x,y: ", x, *y)
	*y += 1
	fmt.Println("x,y: ", x, *y) // nilai x ikut berubah
	tambah(&x, 10)
	fmt.Println("x,y: ", x, *y)
}

// kapan pointer harus digunakan?
// saat harus menyimpan banyak data dengan nilai yang sama
// tujuannya agar hemat memory
