package main

import "fmt"

func mul(numbers ...int) int {
	result := 1
	for _, n := range numbers {
		result *= n
	}
	return result
}

func main() {
	fmt.Println("1x2x..x10 =", mul(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
