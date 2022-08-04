package main

import (
	"fmt"
	"reflect"
)

func random() interface{} {
	return "Hai"
}

func main() {
	val := random()
	fmt.Println("tipe: ", reflect.TypeOf(val))
	switch v := val.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("integer", v)
	default:
		fmt.Println("Unknow.")
	}
}
