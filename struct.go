package main

import "fmt"

// struct
type Manusia struct {
	Nama string
	Umur int
	Hobi [3]string
}

// struct method
func (manusia Manusia) PrintAll() {
	fmt.Println("Nama: ", manusia.Nama)
	fmt.Println("Umur: ", manusia.Umur)
	fmt.Println("Hobi: ", manusia.Hobi)
}

func main() {
	var saya Manusia
	saya.Nama = "Hasban"
	saya.Umur = 16
	saya.Hobi = [3]string{"Coding", "Nonton Youtube", "Bersepeda"}
	saya.PrintAll()

	fmt.Println("===============================================")

	teman := Manusia{
		Nama: "Dava",
		Umur: 16,
		Hobi: [3]string{
			"Main Game",
			"Nonton Youtube",
			"Membaca Buku",
		},
	}
	teman.PrintAll()
}
