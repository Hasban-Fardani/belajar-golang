package main

import "fmt"

func main() {
	// contoh membuat type declaration
	type mp_str map[string]string
	type mp_int map[int]int
	type mp_float64 map[float64]float64
	type mp_float32 map[float32]float32

	// contoh pemakaian
	biodata := mp_str{
		"nama":   "Hasban Fardani",
		"hobi":   "Coding",
		"alamat": "Cimahi",
	}

	// for loop dengan tipedata map
	for k, v := range biodata {
		fmt.Println(k, v)
	}

}
