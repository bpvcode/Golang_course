package main

import "fmt"

func main() {
	slc := make([]int, 10)
	slc[9] = 12
	fmt.Println(slc)
	/*
		NOTA:
			Como este slice não esta referenciado a nenhum array, automaticamente é criado um array interno com 10 posições e o slice aponta para esse array
	*/

	slc = make([]int, 10, 20)
	fmt.Println("Novo slice ->", slc)
	/*
		NOTA:
			Neste caso re-atríbuimos uma nova versão ao slice 'slc'
			Vai ser na mesma um slice inteiros '[]int'
			Vai ter 10 elementos
			Contudo o array interno que será criado terá 20 elementos
	*/
	fmt.Println("Tamanho do slice:", len(slc))
	// len() -> length retorna o tamanho do slice
	fmt.Println("Tamanho do array interno:", cap(slc))
	// cap() -> retorna o tamanho do array interno que o slice referencia (Capacidade total)

	slc = append(slc, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	// append() -> Cria uma nova estrutura sem alterar o array interno. Adiciona ao slice os elementos extra, face aos que já existiam
	fmt.Println("Após append:", slc)
	//COM ESTE append() ATINGIMOS A CAPACIDADE MÁXIMA DO ARRAY INTERNO
	fmt.Println("LIMITE: Tamanho do slice:", len(slc))
	fmt.Println("LIMITE: Tamanho do array interno:", cap(slc))

	slc = append(slc, 9000)
	fmt.Println("Após capacidade máxima atingida:", slc)
	fmt.Println("Tamanho do slice:", len(slc))
	fmt.Println("MAIS CAPACIDADE: Tamanho do array interno:", cap(slc))
	/*
		A capacidade do array interno cresceu
		O slice vai crescendo internamente e referênciando arrays diferentes quando a capacidade máxima do array interno for atingida
	*/
}
