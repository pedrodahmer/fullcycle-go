package main

import "fmt"

type Conta struct {
	saldo int
}

// A SIMULAÇÃO ALTERA APENAS A CÓPIA DO VALOR SALDO
func (c Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

// A FUNÇÃO ALTERAR ALTERA O VALOR NO ENDEREÇO DE MEMÓRIA DO SALDO
func (c *Conta) alterar(valor int) int {
	c.saldo += valor
	return c.saldo
}

func NovaConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	conta := Conta{
		saldo: 123,
	}
	// conta.simular(7)
	conta.alterar(7)
	fmt.Printf("O valor do saldo foi alterado para: %d \n", conta.saldo)

	// CRIANDO UMA NOVA CONTA E ALTERANDO O VALOR DO SALDO
	novaConta := NovaConta()
	fmt.Printf("Saldo da nova conta: %d \n", novaConta.saldo)
	novaConta.alterar(7)
	fmt.Printf("Novo saldo da conta: %d", novaConta.saldo)
}
