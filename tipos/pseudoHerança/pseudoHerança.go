package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //So tem o tipo nao tem nome (tipo carro) - Campo anónimo do tipo carro. Tudo o que é de carro fica disponivel dentro de ferrari
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	/*
		Conseguimos utilizar diretamente os atributos da class carro, a partir de ferrari
		Isto acontece pois tem os campos anónimos ou seja uma referência para a class carro
		Mistura entre composição e herança em Go - ferrari is a carro
	*/
	f.turboLigado = true

	fmt.Printf("O ferrari %s está com o turbo ligado (%t), à velocidade de %d", f.nome, f.turboLigado, f.velocidadeAtual)
}
