// https://tlx.toki.id/courses/basic/chapters/05/problems/F
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
		if number > 0 {
			ceil = floor + 1
		} else {
			ceil = floor
			floor -= 1
		}
		fmt.Println(floor, ceil)
	}
}
