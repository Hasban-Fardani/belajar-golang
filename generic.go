package main

import "fmt"

func normalSum(n ...int) int {
	var total int
	for i := range n {
		total += i
	}
	return total
}

func genericSum[V int | float32 | float64](n ...V) V {
	var total V
	for _, e := range n {
		total += e
	}
	return total
}

func main() {
	fmt.Println("normal sum:", normalSum(1, 2, 3, 4, 5, 6, 7))
	// fmt.Println("normal sum:", normalSum(1.2, 2.3, 6.4, 7.2))  [Error]
	fmt.Println("generic sum:", genericSum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println("generic sum float:", genericSum(1.5, 2.1, 3.5, 4.2, .5, 6.3))

}
