package main

import "fmt"

func main() {
	fmt.Println("Halo")
	fmt.Println("Ini adalah string"[0]) // output: I
	// "string"[n] artinya
	// mengambil karakter dengan index n
	// setiap index dimulai dari 0
	fmt.Println("youtube.com"[0:7]) // output: youtube
	// "string"[n:m] artinya:
	// mengambil string dari n sampai m
	fmt.Println("panjang dari 'Halo' adalah:", len("Halo"))
	// fungsi len => untuk mengeluarkan panjang dari sebuah string
}
