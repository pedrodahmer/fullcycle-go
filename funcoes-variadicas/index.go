package main

import "fmt"

func main() {
	fmt.Println((soma(1, 3, 4, 5, 7, 7)))
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
