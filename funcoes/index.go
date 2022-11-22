package main

import (
	"errors"
	"fmt"
)

func main() {
	// RECEBENDO OS RESULTADOS DA FUNÇÃO DE SOMA
	valor, err := soma(20, 40)
	// CASO TENHA ERRO, MOSTRA O ERRO E NÃO MOSTRA O RESULTADO
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(valor)
}

// FUNÇÃO SIMLES, RETORNANDO APENAS UM VALOR
// func soma(a, b int) {
// 	return a + b
// }

// FUNÇÃO COM DOIS RETORNOS E RETORNANDO ERRO NA CONDICIONAL
func soma(a int, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
