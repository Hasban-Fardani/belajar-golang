package main

import "fmt"

func main() {
	// string
	fmt.Println("string Gits", "Gits")
	// panjang dari string bebas

	// numeric (angka)
	fmt.Println("interger 14: ", 14)
	//          min - max
	// int8  -> -128 - 127
	// int16 -> -32768 - 32767
	// int32 -> -2147483648 - 2147483647
	// int64 -> -9223372036854775808 - 9223372036854775807
	// int == int64
	// sumber: https://gosamples.dev/int-min-max/#:~:text=To%20get%20the%20maximum%20and,is%209223372036854775807%2C%20use%20the%20math.
	fmt.Println("float 1.8: ", 1.8)
	// brp min - max float? kunjungi link berikut:
	// https://gosamples.dev/float64-min-max/#:~:text=%F0%9F%93%8A%20The%20maximum%20and%20minimum%20value%20of%20the%20float%20types%20in%20Go&text=MaxFloat64%20constant.,SmallestNonzeroFloat64%20constant.

	// boolean
	fmt.Println("Benar: ", true)
	fmt.Println("Salah: ", false)
}
