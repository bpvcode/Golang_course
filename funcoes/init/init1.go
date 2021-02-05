package main

import "fmt"

/*
	Função init:
		* É executada antes de ser executada a função main, por mais que nao seja chamada
		* Quando um arquivo é lido ele lê logo a função init() para fazer algum tipo de inicialização

*/

func init() {
	fmt.Println("Inicializando... init()")
}

func main() {
	fmt.Println("MAIN")
}

/*
	AO SER EXECUTADO PELO CODE RUNNER, só tem em conta este ficheiro init1.go, pelo que so imprime inicializando... init()
	Se for pelo terminal vai ter em consideração os dois ficheiro init1.go e init2.go, pelo que imprime os dois, pois ambos fazem parte do mesmo package, main
*/
