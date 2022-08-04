package main

import (
	"errors"
	"fmt"
)

func bagi(angka, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0.")
	} else {
		return angka / pembagi, nil
	}
}
func main() {
	angka := 0
	pembagi := 0
	fmt.Print("masukkan angka: ")
	fmt.Scanln(&angka)
	fmt.Print("masukkan pembagi: ")
	fmt.Scanln(&pembagi)
	hasil, err := bagi(angka, pembagi)
	if err == nil {
		fmt.Println("hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
