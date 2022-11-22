package main

import "fmt"

var b = 1

var c bool

type ID int

func main() {
	c = true
	var d string = "x"

	e := "pedrao"
	e = "pedroca"

	var f ID = 1

	g := "pedrao"

	println(a)
	println(c)
	println(b)
	println(e)
	println(d)
	println(f)

	fmt.Printf("O tipo de f é %T", g)

	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d\n", i, v)
	}
}
