package main

import (
	"fmt"
	_ "image"
)

func main() {
	var a, b, add, product int

	a = 10
	b = 20

	add, _ = addAndProduct(a, b)

	fmt.Println(add, ", ", product)

	fmt.Println(variableReturn(a, b))
}

func addAndProduct(a int, b int) (int, int) {
	return a + b, a * b
}

func variableReturn(a int, b int) (result int) {
	result = (a + b) * (a * b)

	return
}
