package main

import "fmt"

/*
	Struct é uma espécie de classe do java mas em Go
*/
type produto struct { //Declaração de uma struct
	nome     string
	preço    float64
	desconto float64
}

/*
	Receiver - A função será aplicada a um recetor (ao struct)
	É como um método da classe
	O dono da função é a estrutura
*/
func (p produto) precoComDesconto() float64 {
	//Função que tem um receiver (Neste caso o receiver é o produto). Dentro da função, aceder à struct produto tem de ser a partir da letra 'p'.
	// Não precisamos de passar argumentos à funções pois a estrutura ja tem acesso às informações necessária para refletir o preço com desconto (preço, desconto) através do struct
	return p.preço * (1 - p.desconto)
}

func main() {
	//var produto1 produto -> Se declara-se primeiro a variável, depois atribuída sem ':=' mas sim '='
	produto1 := produto{
		nome:     "Bruno",
		preço:    1.79,
		desconto: 0.05,
	}
	fmt.Println("Produto1:", produto1)
	fmt.Println("Produto1 com desconto:", produto1.precoComDesconto())

	produto2 := produto{"Rui", 1000, 0.5}
	fmt.Println("Produto2:", produto2)
	fmt.Println("Produto2 com desconto:", produto2.precoComDesconto())
}

/*
	CONCLUSÃO:
	* Conseguimos definir estruturas (struct) que recebem estado (propriedades)
	* Podemos criar funções que operam só sobre os dados da struct, especificas para a struct (tipo métodos de uma class)
	* Não existe herança em go, mas existe interfaces e polimorfismo
*/
