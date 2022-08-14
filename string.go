package main

import (
	"fmt"
	"strings"
)

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

	// https://dasarpemrogramangolang.novalagung.com/A-strings.html
	// https://pkg.go.dev/strings
	myString := "Hai, my name is Hasban Fardani and I'am 16 years old"
	isExist := strings.Contains(myString, "name")
	fmt.Println(isExist)
	countA := strings.Count(myString, "a")
	fmt.Println("count a:", countA)
	indexHa := strings.Index(myString, "Ha")
	fmt.Println("index ha:", indexHa)

}
