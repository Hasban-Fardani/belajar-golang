package mypackage

import "fmt"

func init(){
	fmt.Println("imported package mypackage")
}

func Tambah(a, b int) int {
	return a + b
}

func Kurang(a, b int) int {  // Wajib menggunakan uppercase di awal function
	// kalo gk pake uppercase gk bisa di import
	return a - b
}