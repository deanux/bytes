package main

import (
	"bytes"
	"fmt"
	"strings"
)

func dean() {
	fmt.Println("hey")
}

func main() {
	testados := ("sea, dog, salt, earth, another, sample, wordw1th, number, go")
	values := []byte("sea, dog, salt, earth, another, sample, wordw1th, number, go")
	alpha := strings.Count(testados, "e")
	result := bytes.Index(values, []byte("dog"))
	fmt.Println(result, alpha)
	dean()
}
