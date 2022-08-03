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
