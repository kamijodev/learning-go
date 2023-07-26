package main

import "fmt"

const x = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"
	var c uint16 = x
	var d float64 = x

	// var b float64 = x

	fmt.Println(x)
	fmt.Println(y)

	// x = x + 1
	// y = "bye"

	fmt.Println(c)
	fmt.Println(d)
	// fmt.Println(y)
}
