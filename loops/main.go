package main

func main() {
	for i := 0; i < 10; i++ {
		println("pedrao")
	}
}

// COMANDO PARA COMPILAR -> go build main.go
// CONFIGURANDO AMBIENTE
// GOOS -> dita o sistema operacional (windows, linux, darwin)
// GOARCH -> dita a arquitetura do sistema (amd64, arm)
// go tool dist list -> mostra todas as opções de distribuições que podem ser usadas (windows, arm)
