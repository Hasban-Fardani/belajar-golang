package main

import (
	"fmt"
)

type HasName interface {
	GetName() string
}

func SayHello(has_name HasName) {
	fmt.Println("Halo", has_name.GetName())
}

// penggunaan infterface kosong
func test(param interface{}) interface{} {
	if param == 1 {
		return "angka 1"
	} else if param == 2 {
		return true
	} else {
		return 0
	}
}

type Orang struct {
	nama string
}

func (orang Orang) GetName() string {
	return orang.nama
}

func main() {
	hasban := Orang{nama: "Hasban"}
	SayHello(hasban)
}
