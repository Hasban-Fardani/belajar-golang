package main

import "fmt"

func main() {
	// membuat tipe data map
	// cara 1
	var saya map[string]string
	saya["nama"] = "Hasban"

	// cara 2
	me := map[string]string{
		"nama":    "Hasban Fardani",
		"sekolah": "SMKN 11 Bandung",
	}

	fmt.Println(me)
	fmt.Println(me["nama"])
	fmt.Println(me["sekolah"])
	fmt.Println(me["hobi"]) // tidak akan mengeluarkan apapun
}
