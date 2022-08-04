package main

import "fmt"

type map_str map[string]string

func CreateMap(name string) map_str {
	if name == "" {
		return nil
	} else {
		return map_str{"name": name}
	}
}

func CreateMap2(name string) map_str {
	return map_str{"name": name}
}

func main() {
	var person = CreateMap("Hasban")
	var person2 = CreateMap2("")

	if person == nil { // singkatnya nil hanya untuk mempersingkat syntax
		fmt.Println("Kosong")
	} else {
		fmt.Println(person)
	}

	if person2["name"] == "" {
		fmt.Println("Kosong")
	} else {
		fmt.Println(person2)
	}
}
