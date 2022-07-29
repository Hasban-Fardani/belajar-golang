package main

import "fmt"

func sapa() {
	// function tanpa parameter dan return value
	fmt.Println("Selamat datang")
}

func halo(nama string) {
	// function dengan parameter tanpa return value
	fmt.Println("Halo", nama)
}

func luas_segitiga(a int, t int) int {
	// funtion dengan parameter dan return value
	return a * t / 2
}

func swap(a int, b int) (int, int) {
	// function multi return
	return b, a
}

func fpb(a int, b int) int {
	// function recrusif
	if b == 0 {
		return a
	} else {
		return fpb(b, a%b)
	}
}

func kpk(a int, b int) int {
	return a * b / fpb(a, b)
}

func main() {
	var (
		nama string
		a    int
		b    int
	)
	fmt.Print("masukkan nama: ")
	fmt.Scanln(&nama)
	halo(nama)

	fmt.Print("masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("masukkan nilai b: ")
	fmt.Scanln(&b)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("=== swaping ===")
	a, b = swap(a, b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("=== FPB dan KPK ===")
	fmt.Println("FPB:", fpb(a, b))
	fmt.Println("KPK:", kpk(a, b))
	/*
		kotretan
			a = 20				a = 28
			b = 15				b = 20
			kpk = 60			kpk = 140
			fpb = 5				fpb = 4
			kpk x fpb = 300		kpk x fpb = 560
			a x b = 300			a x b = 560

			jadi, kpk x fpb = a x b  (berlaku untuk semua angka)
	*/
}
