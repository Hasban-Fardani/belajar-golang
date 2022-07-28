package main

import "fmt"

func abs(number int) int {
	if number < 0 {
		return -number
	} else {
		return number
	}
}

func main() {
	var (
		x1 int
		y1 int
		x2 int
		y2 int
	)
	fmt.Scanln(&x1, &y1, &x2, &y2)
	fmt.Println(abs(x1-x2) + abs(y1-y2))
}
