package main

import "fmt"

type Siswa struct {
	Nama        string
	Kelas       int
	Abjad_kelas string
}

// tanpa pointer
func (siswa Siswa) naik_kelas() {
	fmt.Println(siswa.Nama, " naik kelas tanpa pointer")
	siswa.Kelas += 1
}

func (siswa *Siswa) naik_kelas_beneran() {
	fmt.Println(siswa.Nama, " naik kelas dengan pointer")
	siswa.Kelas += 1
}

func main() {
	siswa1 := Siswa{"Hasban", 9, "a"}
	fmt.Println(siswa1)
	siswa1.naik_kelas()
	fmt.Println(siswa1) // tidak akan berubah

	fmt.Println()

	siswa2 := &Siswa{"Dava", 9, "b"}
	fmt.Println(*siswa2)
	siswa2.naik_kelas_beneran()
	fmt.Println(*siswa2)
}
