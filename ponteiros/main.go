package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10

	// MOSTRANDO O VALOR PRESENTE NO ENDEREÇO RESERVADO PARA A
	println(a)

	// MOSTRANDO O ENDEREÇO RESERVADO PARA A
	println(&a)

	// DEFININDO UM VARIÁVEL PARA CONTER O VALOR DO ENDEREÇO DE A
	var ponteiro *int = &a
	println(ponteiro)

	// ALTERANDO O VALOR DE A USANDO O PONTEIRO
	*ponteiro = 20

	println(a)

	// CRIANDO UM PONTEIRO DE FORMA SIMPLIFICADA
	b := &a
	println(b)

	// MOSTRANDO O VALOR QUE ESTÁ GUARDADO NO ENDEREÇO
	// QUE b REPRESENTA
	println(*b)

	// ALTERANDO O VALOR DE a ATRAVÉS DE b
	*b = 30
	println(a)
}
