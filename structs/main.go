package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func main() {
	pedrao := Cliente{
		Nome:  "Pedrao",
		Idade: 21,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", pedrao.Nome, pedrao.Idade, pedrao.Ativo)

	pedrao.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", pedrao.Nome, pedrao.Idade, pedrao.Ativo)

	pedrao.Endereco = Endereco{
		Logradouro: "Gonçalves Ledo",
		Numero:     "254",
		Cidade:     "Porto Alegre",
		Estado:     "Rio Grande do Sul",
	}

	pedrao.Desativar()

	minhaEmpresa := Empresa{
		Nome: "Empresa do pedrao",
	}

	Desativacao(pedrao)
	Desativacao(minhaEmpresa)
}
