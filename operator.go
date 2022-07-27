package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	fmt.Print("masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("masukkan nilai b: ")
	fmt.Scanln(&b)

	// operator matematika
	fmt.Println("\n===== Operator Matematika =====")
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a + b = ", a+b)
	fmt.Println("a % b = ", a%b)

	// operator perbandingan
	fmt.Println("\n===== Operator Perbandingan =====")
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b :", a > b)
	fmt.Println("a < b :", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// operator boolean
	var (
		x bool
		y bool
	)
	fmt.Print("\nmasukkan nilai x [boolean]: ")
	fmt.Scanln(&x)
	fmt.Print("masukkan nilai y [boolean]: ")
	fmt.Scanln(&y)
	fmt.Println("\n===== Operator boolean =====")
	fmt.Println("  x   :", x)
	fmt.Println("  y   :", y)
	fmt.Println(" !x   :", !x)
	fmt.Println(" !y   :", !y)
	fmt.Println("x && y:", x && y)
	fmt.Println("x || y:", x || y)
}
