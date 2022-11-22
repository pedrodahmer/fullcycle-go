package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// PRINTANDO O COMPRIMENTO, CAPACIDADE E VALORES DO SLICE
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// DEIXANDO DE FORA TUDO QUE VEM APÓS O ITEM 0
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// DEIXANDO DE FORA TUDO QUE VEM APÓS O ITEM 4
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// DEIXANDO DE FORA OS PRIMEIROS 2 ITENS
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// AUMENTANDO A CAPACIDADE DO SLICE
	// O APPEN CRIA UM ARRAY MAIOR POR BAIXO DOS PANOS
	// COM O DOBRO DA CAPACIDADE INICIAL
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
