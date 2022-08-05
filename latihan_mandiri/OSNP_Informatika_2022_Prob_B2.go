package main

import "fmt"

func main() {
	var (
		A    int
		B    int
		temp int
	)
	fmt.Scanln(&A, &B)
	for B != 0 {
		temp = A
		A = B
		B = temp % B
	}
	fmt.Println(A)
}
