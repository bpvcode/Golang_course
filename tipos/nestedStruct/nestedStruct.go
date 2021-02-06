package main

import "fmt"

//STRUCT DENTRO DE STRUCT

type item struct {
	id         int
	quantidade int
	preço      float64
}

type pedido struct {
	idUser int
	itens  []item //Dentro desta struct temos outra struct de item. Ou seja, o pedido do cliente pode ter vários itens
}

func (p pedido) valorTotalDoPedido() float64 {
	valorTotal := 0.0

	for _, item := range p.itens {
		valorTotal += item.preço * (float64(item.quantidade))
	}
	return valorTotal
}

func main() {
	item1 := item{
		id:         1,
		quantidade: 3,
		preço:      9.99,
	}
	item2 := item{2, 1, 20.0}

	pedido1 := pedido{
		idUser: 1,
		itens: []item{
			item1,
			item2,
		},
	}

	fmt.Printf("VALOR TOTAL DO PEDIDO1: %.2f€", pedido1.valorTotalDoPedido())
}
