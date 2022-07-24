package main // import package

import "fmt" // import fmt dari package

func main() {
	fmt.Println("Halo dunia!") // cetak 'Halo dunia' + enter text
	fmt.Print("halo")          // cetak 'halo' tanpa enter
	fmt.Print("indonesia\n")   // cetak 'indonesia' + enter ('\n' adalah karakter spesial)
	fmt.Printf("hari ini tanggal %d bulan %s", 23, "Juli")
	// %d akan di ganti dengan 23, dan %s akan diganti dengan "Juli"
}
