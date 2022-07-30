package main

import "fmt"

func fpb(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return fpb(b, a%b)
	}
}

func main() {
	var (
		N int
		A int
		B int
	)
	fmt.Scanln(&N, &A, &B)
	kpk := A * B / fpb(A, B)
	fmt.Println(kpk/A + kpk/B)
}
