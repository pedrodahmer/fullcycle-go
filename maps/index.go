package main

import "fmt"

func main() {
	// CRIANDO O MAP SALARIOS
	salarios := map[string]int{"Pedrao": 2500, "Max": 1000}
	// MOSTRANDO APENAS O PEDRAO
	fmt.Println(salarios["Pedrao"])

	// DELETANDO UM ITEM
	delete(salarios, "Pedrao")

	// ADICIONANDO UM ITEM
	salarios["Marcos"] = 17
	fmt.Println(salarios["Marcos"])

	// CRIANDO UM MAP USANDO MAKE
	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	// PERCORRENDO O SALÁRIO DE TODO MUNDO
	for nome, salario := range salarios {
		fmt.Printf("O salário do %s é %d\n", nome, salario)
	}

	// SEM O NOME (USANDO O BLANK IDENTIFIER)
	for _, salario := range salarios {
		fmt.Printf("%d\n", salario)
	}
}
