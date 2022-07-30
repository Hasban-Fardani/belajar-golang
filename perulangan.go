package main

import "fmt"

func main() {
	fmt.Println("===== Perulangan =====")
	var (
		a int
		b int
	)
	fmt.Print("masukkan angka: ")
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		fmt.Printf("\nperulangan ke-%d", i)
	}
	fmt.Println("\n======================")

	for a != 0 {
		fmt.Println("ini infinite loop selama a!=0")
		fmt.Println("loop ke", b)
		b += 1
		fmt.Scanln(&a)
	}

	test := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range test {
		fmt.Println(i)
	}

	me := map[string]string{
		"nama": "Hasban Fardani",
		"Hobi": "Coding",
	}
	for key, value := range me {
		fmt.Println(key, value)
	}
}
