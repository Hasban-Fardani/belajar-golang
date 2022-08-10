package main

import (
	"fmt"
	"time"
)

func print(banyak int, txt interface{}) {
	for i := 0; i < banyak; i++ {
		fmt.Println(txt)
	}
}


func main() {

	// go print(5, "apa kabar")
	// go print(3, "Halo halo")
	// print(10, "Hai semua")

	time.Sleep(1 * time.Second)
}
