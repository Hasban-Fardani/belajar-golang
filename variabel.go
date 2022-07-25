package main

import "fmt"

func main() {
	// cara 1
	var nama string
	// var: untuk mendefinisikan variabel
	// nama: nama variabel
	// string: tipe date variabel
	nama = "Hasban Fardani"
	// assign "Hasban Fardani" ke variabel nama
	fmt.Println("nama:", nama) // cetak variabel nama
	//cara 2
	var friend = "Dava Ardiansyah"
	fmt.Println("my friend:", friend)

	var no_absen int8 = 13 // memaksakan 13 bertipedata int8
	fmt.Println("no absenku:", no_absen)

	// cara 3
	umur := 15 // bakal error kalo dikasih tipedata didepan nama variabel
	// umur int := 15  (bakal error)
	fmt.Println("umur:", umur, "(tahun lalu)")
	// ganti nilai variabel umur
	umur = 16
	fmt.Println("umur sekarang:", umur)

	// cara 4
	var (
		hobi1 = "nonton youtube"
		hobi2 = "ngoding"
		hobi3 = "baca komik"
		value = 95 // bisa berbagai tipe dalam sekali deklarai var
	)
	fmt.Println("hobiku:", hobi1, ",", hobi2, ",", hobi3)
	fmt.Println("value:", value)
}
