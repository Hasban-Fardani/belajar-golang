package main

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()

	trueResult := "Hello Hasban"
	result := sayHello("Hasban")
	t.Logf("result: %s", result)
	if result != "Hello Hasban" {
		panic("Error, expeted " + trueResult + " not " + result)
	}
}
