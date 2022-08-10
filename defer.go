package main

import "fmt"

func main(){
	defer fmt.Println("Jumpa lagi")  // akan dieksekusi di akhir
	fmt.Println("Selamat datang")
	fmt.Println("Senang bertemu denganmu")
}