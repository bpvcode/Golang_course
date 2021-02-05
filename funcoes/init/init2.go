package main

import "fmt"

func init() {
	fmt.Println("Inicializando... init() - 2")
}

/*
	COMO DENTRO DO PACKAGE main JÁ TEMOS UMA FUNÇÃO main(), NO FICHEIRO init.go
	RETIRAMOS A FUNÇÃO main() DESTE FICHEIRO E SERÁ LIDA NA MESMA NO OUTRO FICHEIRO, ANTES DA FUNÇÃO main() SER INICIADA

	SÓ SERÁ EXECUTADO PELO TERMINAL (nao pelo code runner)
*/
