package main

import "fmt"

func main() {
	var number float64
	var ceil int

	fmt.Scanln(&number)
	floor := int(number)
	comma := number - float64(floor)

	if comma == 0. {
		fmt.Println(number, number)
	} else {
		ceil = floor + 1
		fmt.Println(floor, ceil)
	}
}
